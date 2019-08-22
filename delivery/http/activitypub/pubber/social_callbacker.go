package pubber

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
)

// SocialCallbacker provides an Application hooks into the lifecycle of the
// ActivityPub processes for both client-to-server and server-to-server
// interactions. These callbacks are called after their spec-compliant actions
// are completed, but before inbox forwarding and before delivery.
//
// Note that at minimum, for inbox forwarding to work correctly, these
// Activities must be stored in the client application as a system of record.
//
// Note that modifying the ActivityStream objects in a callback may cause
// unintentionally non-standard behavior if modifying core attributes, but
// otherwise affords clients powerful flexibility. Use responsibly.
type SocialCallbacker struct {
	db *sql.DB
}

// Create Activity callback.
func (m *SocialCallbacker) Create(c context.Context, s *streams.Create) error {
	log.Println("SocialCallbacker CREATE", s)
	v := s.Raw()
	oc := v

	for i := 0; i < oc.ObjectLen(); i++ {
		obj := oc.GetObject(i)
		switch v := obj.(type) {
		case vocab.DocumentType:
			query := "INSERT INTO documents (url, name, author, created_at) VALUES ($1, $2, $3, $4);"
			url := obj.GetUrlAnyURI(0).String()
			title := obj.GetNameString(0)
			postAuthor := obj.GetAttributedToIRI(0).String()
			_, err := m.db.Exec(query, url, title, postAuthor, time.Now())
			if err != nil {
				return fmt.Errorf("error in m.db.Exec:%s", err)
			}
		default:
			log.Println("UNKNOWN Type %T", v)
		}
	}
	return nil
}

// Update Activity callback.
func (m *SocialCallbacker) Update(c context.Context, s *streams.Update) error {
	log.Println("SocialCallbacker UPDATE", s)
	return nil
}

// Delete Activity callback.
func (m *SocialCallbacker) Delete(c context.Context, s *streams.Delete) error {
	log.Println("SocialCallbacker DELETE", s)
	return nil
}

// Add Activity callback.
func (m *SocialCallbacker) Add(c context.Context, s *streams.Add) error {
	log.Println("SocialCallbacker ADD", s)
	return nil
}

// Remove Activity callback.
func (m *SocialCallbacker) Remove(c context.Context, s *streams.Remove) error {
	log.Println("SocialCallbacker REMOVE", s)
	return nil
}

// Like Activity callback.
func (m *SocialCallbacker) Like(c context.Context, s *streams.Like) error {
	log.Println("SocialCallbacker LIKE", s)
	v := s.Raw()

	query := "INSERT INTO liked (uri, author) VALUES ($1, $2);"
	url := v.GetObjectIRI(0).String()
	author := v.GetActorIRI(0).String()
	_, err := m.db.Exec(query, url, author)
	if err != nil {
		return fmt.Errorf("error in m.db.Exec:%s", err)
	}

	return nil
}

// Block Activity callback. By default, this implmentation does not
// dictate how blocking should be implemented, so it is up to the
// application to enforce this by implementing the FederateApp
// interface.
func (m *SocialCallbacker) Block(c context.Context, s *streams.Block) error {
	log.Println("SocialCallbacker BLOCK", s)
	return nil
}

// Follow Activity callback. In the special case of server-to-server
// delivery of a Follow activity, this implementation supports the
// option of automatically replying with an 'Accept', 'Reject', or
// waiting for human interaction as provided in the FederateApp
// interface.
//
// In the special case that the FederateApp returned AutomaticAccept,
// this library automatically handles adding the 'actor' to the
// 'followers' collection of the 'object'.
func (m *SocialCallbacker) Follow(c context.Context, s *streams.Follow) error {
	log.Println("SocialCallbacker FOLLOW", s)
	return nil
}

// Undo Activity callback. It is up to the client to provide support
// for all 'Undo' operations; this implementation does not attempt to
// provide a generic implementation.
func (m *SocialCallbacker) Undo(c context.Context, s *streams.Undo) error {
	log.Println("SocialCallbacker UNDO", s)
	return nil
}

// Accept Activity callback. In the special case that this 'Accept'
// activity has an 'object' of 'Follow' type, then the library will
// handle adding the 'actor' to the 'following' collection of the
// original 'actor' who requested the 'Follow'.
func (m *SocialCallbacker) Accept(c context.Context, s *streams.Accept) error {
	log.Println("SocialCallbacker ACCEPT", s)
	return nil
}

// Reject Activity callback. Note that in the special case that this
// 'Reject' activity has an 'object' of 'Follow' type, then the client
// MUST NOT add the 'actor' to the 'following' collection of the
// original 'actor' who requested the 'Follow'.
func (m *SocialCallbacker) Reject(c context.Context, s *streams.Reject) error {
	log.Println("SocialCallbacker REJECT", s)
	return nil
}
