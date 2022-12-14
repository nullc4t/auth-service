package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	stdjwt "github.com/golang-jwt/jwt/v4"
	"github.com/nullc4t/auth-service/pkg/access"
	"github.com/nullc4t/auth-service/pkg/account"
	"github.com/nullc4t/auth-service/pkg/auth"
	"github.com/nullc4t/auth-service/pkg/jwt"
	"github.com/nullc4t/auth-service/pkg/logger"
	"github.com/nullc4t/auth-service/pkg/mgmt"
	"github.com/nullc4t/auth-service/pkg/permission"
	svcsrv "github.com/nullc4t/auth-service/pkg/service"
	"github.com/nullc4t/auth-service/pkg/types"
	"github.com/nullc4t/auth-service/pkg/user"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type testSuite struct {
	suite.Suite
	db   *gorm.DB
	acco account.Repo
	user user.Service
	mgmt mgmt.Service
	svc  svcsrv.Service
	auth auth.Service
	p    permission.Service
	jwt  jwt.Service
}

func (s *testSuite) SetupSuite() {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		"localhost",
		"postgres",
		"password",
		"postgres",
		5432,
		"Europe/Kiev",
	)))
	s.Require().NoError(err)
	s.db = db
	s.acco = account.NewLoggingMiddleware(
		logger.New("[ account auth ]\t"),
		account.New(db),
	)
	s.user = user.NewLoggingMiddleware(logger.New("[ user ]\t"), user.New(db))
	s.p = permission.New(db)

	s.svc = svcsrv.New(logger.New("[ auth ]\t"), db)
	s.mgmt = mgmt.New(logger.New("[ mgmt ]\t"), db, s.svc, s.acco, s.p, s.user)

	var privateKey = make([]byte, 64)
	_, err = rand.Read(privateKey)
	s.Require().NoError(err)

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	s.Require().NoError(err)

	s.jwt = jwt.New(
		logger.New("[ jwt ]\t"),
		stdjwt.SigningMethodES256,
		[]string{stdjwt.SigningMethodES256.Name},
		jwt.AccessClaimsFactory,
		key,
		&key.PublicKey,
	)

	s.auth = auth.New(logger.New("[ auth ]\t"), s.mgmt, s.jwt)

	tables, err := s.db.Migrator().GetTables()
	s.Require().NoError(err)
	for _, table := range tables {
		s.Require().NoError(s.db.Debug().Migrator().DropTable(table))
	}
	s.Require().NoError(s.db.Debug().AutoMigrate(
		&types.User{}, &types.Account{}, &types.Service{}, &types.Permission{},
	))
}

func (s *testSuite) TestAuth() {
	az := access.NewHelperFromPermissions("read", "write")
	var svc types.Service
	s.Run("create auth", func() {
		sv, err := s.svc.Create(context.TODO(), "test")
		s.Require().NoError(err)
		svc = *sv
	})
	var perms []*types.Permission
	for name, access := range az.ByName() {
		s.Run(fmt.Sprint("create permission", name, access, "auth:", svc.Name, svc.ID), func() {
			v, err := s.mgmt.CreatePermission(context.TODO(), svc.ID, name, &access)
			s.Require().NoError(err)
			s.T().Log(v.ID, v.Name, v.Access, v.ServiceID)
			perms = append(perms, v)
		})
	}
	svcNames := []string{"", svc.Name}
	accIDs := []uint32{0, 1}
	var j int
	for i := 0; i < 2; i++ {
		for _, name := range svcNames {
			for _, accID := range accIDs {
				j++
				s.Run(fmt.Sprint("register", i, name, accID), func() {
					ok, err := s.auth.Register(context.TODO(), fmt.Sprint("test", j), "test", name, accID)
					s.Require().NoError(err)
					s.Require().True(ok)
				})

				for _, perm := range perms {
					s.Run("set user permissions", func() {
						u, err := s.user.Get(context.TODO(), &types.User{Name: fmt.Sprint("test", j)})
						s.Require().NoError(err)
						s.T().Log(u.ID, u.AccountID, u.Name, u.Permissions)
						v, err := s.mgmt.AddUserPermission(context.TODO(), &types.Permission{Model: types.Model{ID: perm.ID}}, u.ID)
						s.Require().NoError(err)
						s.Require().True(v)
						u, err = s.user.Get(context.TODO(), &types.User{Name: fmt.Sprint("test", j)})
						s.Require().NoError(err)
						s.T().Log(u.ID, u.AccountID, u.Name, u.Permissions)
					})
				}

				s.Run(fmt.Sprint("login", i, name, accID), func() {
					t, err := s.auth.Login(context.TODO(), fmt.Sprint("test", j), "test", name)
					s.Require().NoError(err)
					s.T().Log(t.AccessToken)
					claims, err := s.jwt.VerifyAccessToken(t.AccessToken)
					s.Require().NoError(err)
					s.T().Log(claims)
					s.Require().True(az.Access("read", "write").Check(claims.(*jwt.AccessClaims).Access))
				})
			}
		}
	}
}

func (s *testSuite) TestService() {
	az := access.NewHelperFromPermissions("active", "read", "write", "admin", "root")
	var testServices []*types.Service
	for i := 0; i < 5; i++ {
		s.Run(fmt.Sprint("create auth", i+1), func() {
			v, err := s.svc.Create(context.TODO(), fmt.Sprint("auth", i+1))
			s.Require().NoError(err)
			s.T().Logf("%+v", v)
			testServices = append(testServices, v)
		})
	}

	for i, testService := range testServices {
		for name, access := range az.ByName() {
			s.Run(fmt.Sprint(i, testService.Name, "add permission", name), func() {
				v, err := s.mgmt.CreatePermission(context.TODO(), testService.ID, name, &access)
				s.Require().NoError(err)
				s.T().Logf("%d\t%+v", i, v)
			})
		}
	}

	var testPermissions []*types.Permission

	for i, svc := range testServices {
		s.Run(fmt.Sprint(i, "get", svc.Name, "permissions"), func() {
			v, err := s.mgmt.GetFilteredPermissions(context.TODO(), &types.Permission{ServiceID: svc.ID})
			s.Require().NoError(err)
			s.T().Logf("%d\t%+v", i, v)
			testPermissions = append(testPermissions, v...)
		})
	}

	var testUsers []*types.User

	for i := 0; i < 5; i++ {
		s.Run(fmt.Sprint("create user", i+1), func() {
			name := fmt.Sprint("user", i)
			v, err := s.user.CreateWithLoginPassword(context.TODO(), name, name)
			s.Require().NoError(err)
			s.T().Logf("%d\t%+v", i, v)
			testUsers = append(testUsers, v)
		})
	}

	for _, user := range testUsers {
		for _, svc := range testServices {
			p, err := s.mgmt.GetFilteredPermissions(context.TODO(), &types.Permission{ServiceID: svc.ID})
			s.Require().NoError(err)
			for _, permission := range p {
				s.Run(fmt.Sprint("add permission", user.ID, svc.ID, permission.Name), func() {
					v, err := s.mgmt.AddUserPermission(context.TODO(), permission, user.ID)
					s.Require().NoError(err)
					s.T().Logf("%+v", v)
				})
			}
		}
	}

	for name, access := range az.ByName() {
		s.T().Log(name, access)
	}

	for _, user := range testUsers {
		for _, svc := range testServices {
			s.Run(fmt.Sprint("get", user.Name, svc.Name, "permissions"), func() {
				v, err := s.mgmt.GetUserPermissions(context.TODO(), user.ID)
				s.Require().NoError(err)
				s.T().Log("user", user.Name, user.ID)
				for i, permission := range v {
					s.T().Log(i, permission.ID, permission.ServiceID, permission.Name, permission.Access)
				}
			})
		}
	}

	testUsers, err := s.user.GetAll(context.TODO())
	s.Require().NoError(err)
	for i, user := range testUsers {
		s.T().Log(i, user.ID, user.Name, user.Permissions)
		for _, permission := range user.Permissions {
			s.T().Log(permission.ID, permission.ServiceID, permission.Name, permission.Access)
		}
	}
	//for _, user := range testUsers {
	//	s.Run("delete permission", func() {
	//		ok, err := s.authz.RemoveUserPermission(context.TODO(), types.Permission{Name: "root"}, user.ID)
	//		s.Require().NoError(err)
	//		s.Require().True(ok)
	//	})
	//}
	testUsers, err = s.user.GetAll(context.TODO())
	s.Require().NoError(err)
	for i, user := range testUsers {
		s.T().Log(i, user.ID, user.Name, user.Permissions)
		for _, permission := range user.Permissions {
			s.T().Log(permission.ID, permission.ServiceID, permission.Name, permission.Access)
		}
	}
}

func (s *testSuite) TestPermissions() {
	az := access.NewHelperFromPermissions("active", "read", "write", "admin", "root")

	svc, err := s.svc.Create(context.TODO(), "svc1")
	s.Require().NoError(err)

	svc, err = s.svc.Get(context.TODO(), svc)
	s.Require().NoError(err)
	s.T().Log(svc)

	a := az.ByName()["active"]
	p, err := s.mgmt.CreatePermission(context.TODO(), svc.ID, "active", &a)
	s.Require().NoError(err)
	s.T().Log(p)

	ps, err := s.mgmt.GetFilteredPermissions(context.TODO(), &types.Permission{ServiceID: svc.ID})
	s.Require().NoError(err)
	s.T().Log(ps)

}

func TestAuth(t *testing.T) {
	suite.Run(t, new(testSuite))
}
