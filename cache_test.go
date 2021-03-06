package dqrack

import (
	"testing"
)

func TestCacheGet(t *testing.T) {
	var cacheNode Tnode
	jsonCast(&cacheNode, map[string]interface{}{
		"Field1": "hello!",
		"Sub": map[string]string{
			"HiHi": "hi!",
		},
	})

	cacheNode.identity = cacheNode.GetName()
	n0, err := tDq.PutNode(cacheNode)
	if err != nil {
		t.Error(err)
		return
	}

	n1, err := tDq.Node(cacheNode)
	if err != nil {
		t.Error(err)
		return
	}

	if n1.String() != n0.String() {
		t.Errorf("node identities are different, got %s and %s for %s", n1, n0, cacheNode.identity)
	}
}

func TestCacheGet_Miss(t *testing.T) {
	var cacheNode Tnode
	jsonCast(&cacheNode, map[string]interface{}{
		"Field1": "hello!",
		"Sub": map[string]string{
			"HiHi": "hi!",
		},
	})

	cacheNode.identity = cacheNode.GetName()
	n0, err := tDq.PutNode(cacheNode)
	if err != nil {
		t.Error(err)
		return
	}

	tDq.lru.Purge()

	n1, err := tDq.Node(cacheNode)
	if err != nil {
		t.Error(err)
		return
	}

	if n1.String() != n0.String() {
		t.Errorf("node identities are different, got %s and %s for %s", n1, n0, cacheNode.identity)
	}
}

func TestCacheGet_HardMiss(t *testing.T) {
	n, err := tDq.cacheGet("not-good")
	if err != nil {
		t.Error(err)
		return
	}

	if n.String() != "" {
		t.Error("node should be nil")
	}
}
