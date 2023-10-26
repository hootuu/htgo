package htpeer

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/sys"
)

type Peer struct {
	VN         vn.ID
	SP         sp.ID
	PrivateKey ki.PRI
}

func (p *Peer) Verify() *errors.Error {
	if len(p.VN) == 0 {
		return errors.Verify("require peer.VN")
	}
	if len(p.PrivateKey) == 0 {
		return errors.Verify("require peer.PrivateKey")
	}
	return nil
}

var gPeer *Peer

func Initialize(peer *Peer) {
	if peer == nil {
		sys.Error("htpeer.Initialize must set peer, it is null")
		sys.Exit(errors.Verify("require peer"))
		return
	}

	if err := peer.Verify(); err != nil {
		sys.Exit(err)
		return
	}

	gPeer = peer

	sys.Info("# HT Peer Start At VN: ", gPeer.VN)
	sys.Info("# HT Peer Start At SP: ", gPeer.SP)
}

func Here() *Peer {
	if gPeer == nil {
		sys.Warn("Must Call htpeer.Initialize First")
		sys.Exit(errors.Sys("Must Call htpeer.Initialize First"))
		return nil
	}
	return gPeer
}
