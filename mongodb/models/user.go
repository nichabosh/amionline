package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User defines the document structure used to represent an end-user. The
// uuid field acts as the document's unique identifier, whereas the remaining
// fields persist personal information and other relevant/necessary metadata.
type User struct {
	uuid      primitive.ObjectID `bson:"_id,omitempty"`
	createdAt primitive.DateTime `bson:"createdAt,omitempty"`
	updatedAt primitive.DateTime `bson:"updatedAt,omitempty"`

	firstName string `bson:"firstName" required:"true"`
	lastName  string `bson:"lastName" required:"true"`
	username  string `bson:"username" required:"true"`
	email     string `bson:"email" required:"true"`
	password  string `bson:"password" required:"true"`
}

// NewUser takes in a set of user-provided credentials (i.e. first and
// last name, username, email, password), instantiates a new User object
// with unique identifiers and timestamp information, and hashes/salts the
// passed-in, plaintext password to ensure secure storage in the database.
func NewUser(fname, lname, uname, email, pass string) (*User, error) {
	uuid := primitive.NewObjectID()
	now := primitive.NewDateTimeFromTime(time.Now())
	if err := hashAndSaltPassword(&pass); err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	return &User{
		uuid:      uuid,
		createdAt: now,
		updatedAt: now,

		firstName: fname,
		lastName:  lname,
		username:  uname,
		email:     email,
		password:  pass,
	}, nil
}

func hashAndSaltPassword(password *string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(*password), bcrypt.DefaultCost)
	*password = string(hashedBytes)
	return err
}
