package main

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Graph interface {
	UpsertLink(link *Link) error
	FindLink(id ulid.ULID) (*Link, error)
	UpsertEdge(edge *Edge) error
	RemoveStaleEdges(fromID ulid.ULID, updatedBefore time.Time) error
	Links(fromID, toID ulid.ULID, retrievedBefore time.Time) (LinkIterator,
		error)
	Edges(fromID, toID ulid.ULID, updatedBefore time.Time) (EdgeIterator,
		error)
}

type Link struct {
	ID          ulid.ULID
	URL         string
	RetrievedAt time.Time
}

type Edge struct {
	ID        ulid.ULID
	Src       ulid.ULID
	Dst       ulid.ULID
	UpdatedAt time.Time
}

type LinkIterator interface {
	Iterator
	Link() *Link
}

type EdgeIterator interface {
	Iterator
	Edge() *Edge
}

type Iterator interface {
	Next() bool
	Error() error
	Close() error
}
