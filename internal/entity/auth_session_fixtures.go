package entity

import "time"

type SessionMap map[string]Session

func (m SessionMap) Get(name string) Session {
	if result, ok := m[name]; ok {
		return result
	}

	return Session{}
}

func (m SessionMap) Pointer(name string) *Session {
	if result, ok := m[name]; ok {
		return &result
	}

	return &Session{}
}

var SessionFixtures = SessionMap{
	"alice": {
		ID:        "69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac0",
		user:      UserFixtures.Pointer("alice"),
		UserUID:   UserFixtures.Pointer("alice").UserUID,
		UserName:  UserFixtures.Pointer("alice").UserName,
		ExpiresAt: TimeStamp().Add(time.Hour * 168),
	},
	"bob": {
		ID:        "69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac1",
		user:      UserFixtures.Pointer("bob"),
		UserUID:   UserFixtures.Pointer("bob").UserUID,
		UserName:  UserFixtures.Pointer("bob").UserName,
		ExpiresAt: TimeStamp().Add(time.Hour * 168),
	},
	"unauthorized": {
		ID:        "69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac2",
		user:      UserFixtures.Pointer("unauthorized"),
		UserUID:   UserFixtures.Pointer("unauthorized").UserUID,
		UserName:  UserFixtures.Pointer("unauthorized").UserName,
		ExpiresAt: TimeStamp().Add(time.Hour * 168),
	},
	"visitor": {
		ID:       "69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac3",
		user:     &Visitor,
		UserUID:  Visitor.UserUID,
		UserName: Visitor.UserName,
		DataJSON: []byte(`{"tokens":["1jxf3jfn2k"],"shares":["at9lxuqxpogaaba8"]}`),
		data: &SessionData{
			Tokens: []string{"1jxf3jfn2k"},
			Shares: UIDs{"at9lxuqxpogaaba8"},
		},
		ExpiresAt: TimeStamp().Add(time.Hour * 168),
	},
	"friend": {
		ID:        "69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac4",
		user:      UserFixtures.Pointer("friend"),
		UserUID:   UserFixtures.Pointer("friend").UserUID,
		UserName:  UserFixtures.Pointer("friend").UserName,
		ExpiresAt: TimeStamp().Add(time.Hour * 168),
	},
}

// CreateSessionFixtures inserts known entities into the database for testing.
func CreateSessionFixtures() {
	for _, entity := range SessionFixtures {
		Db().Create(&entity)
	}
}
