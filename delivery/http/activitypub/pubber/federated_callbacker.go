package pubber

import (
	"context"
	"log"

	"github.com/go-fed/activity/streams"
)

// FederatedCallbacker provides an Application hooks into the lifecycle of the
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
type FederatedCallbacker struct{}

// Create Activity callback.
func (m *FederatedCallbacker) Create(c context.Context, s *streams.Create) error {
	log.Println("FederatedCallbacker CREATE", s)
	return nil
}

// Update Activity callback.
func (m *FederatedCallbacker) Update(c context.Context, s *streams.Update) error {
	log.Println("FederatedCallbacker UPDATE", s)
	return nil
}

// Delete Activity callback.
func (m *FederatedCallbacker) Delete(c context.Context, s *streams.Delete) error {
	log.Println("FederatedCallbacker DELETE", s)
	return nil
}

// Add Activity callback.
func (m *FederatedCallbacker) Add(c context.Context, s *streams.Add) error {
	log.Println("FederatedCallbacker ADD", s)
	return nil
}

// Remove Activity callback.
func (m *FederatedCallbacker) Remove(c context.Context, s *streams.Remove) error {
	log.Println("FederatedCallbacker REMOVE", s)
	return nil
}

// Like Activity callback.
func (m *FederatedCallbacker) Like(c context.Context, s *streams.Like) error {
	log.Println("FederatedCallbacker LIKE", s)
	return nil
}

// Block Activity callback. By default, this implmentation does not
// dictate how blocking should be implemented, so it is up to the
// application to enforce this by implementing the FederateApp
// interface.
func (m *FederatedCallbacker) Block(c context.Context, s *streams.Block) error {
	log.Println("FederatedCallbacker BLOCK", s)
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
func (m *FederatedCallbacker) Follow(c context.Context, s *streams.Follow) error {
	log.Println("FederatedCallbacker FOLLOW", s)
	return nil
}

// Undo Activity callback. It is up to the client to provide support
// for all 'Undo' operations; this implementation does not attempt to
// provide a generic implementation.
func (m *FederatedCallbacker) Undo(c context.Context, s *streams.Undo) error {
	log.Println("FederatedCallbacker UNDO", s)
	return nil
}

// Accept Activity callback. In the special case that this 'Accept'
// activity has an 'object' of 'Follow' type, then the library will
// handle adding the 'actor' to the 'following' collection of the
// original 'actor' who requested the 'Follow'.
func (m *FederatedCallbacker) Accept(c context.Context, s *streams.Accept) error {
	log.Println("FederatedCallbacker ACCEPT", s)
	return nil
}

// Reject Activity callback. Note that in the special case that this
// 'Reject' activity has an 'object' of 'Follow' type, then the client
// MUST NOT add the 'actor' to the 'following' collection of the
// original 'actor' who requested the 'Follow'.
func (m *FederatedCallbacker) Reject(c context.Context, s *streams.Reject) error {
	log.Println("FederatedCallbacker REJECT", s)
	return nil
}
