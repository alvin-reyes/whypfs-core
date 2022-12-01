package whypfs

import (
	"time"

	datastore "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/multiformats/go-multiaddr"
)

// A list of peer addresses that are used to bootstrap the network
var defaultTestBootstrapPeers []multiaddr.Multiaddr

//"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
//"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
//"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
//"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",

// Creating a list of multiaddresses that are used to bootstrap the network.
func DefaultBootstrapPeers() []peer.AddrInfo {

	for _, s := range []string{
		"/ip4/145.40.90.135/tcp/6746/p2p/12D3KooWNTiHg8eQsTRx8XV7TiJbq3379EgwG6Mo3V3MdwAfThsx", // AR
	} {
		ma, err := multiaddr.NewMultiaddr(s)
		if err != nil {
			panic(err)
		}
		defaultTestBootstrapPeers = append(defaultTestBootstrapPeers, ma)
	}

	peers, _ := peer.AddrInfosFromP2pAddrs(defaultTestBootstrapPeers...)
	return peers
}

// NewInMemoryDatastore provides a sync Datastore that lives in-memory only
// and is not persisted.
func NewInMemoryDatastore() datastore.Batching {
	return dssync.MutexWrap(datastore.NewMapDatastore())
}

var connMgr, _ = connmgr.NewConnManager(100, 600, connmgr.WithGracePeriod(time.Minute))

// Libp2pOptionsExtra provides some useful libp2p options
// to create a fully featured libp2p Host. It can be used with
// SetupLibp2p.
var Libp2pOptionsExtra = []libp2p.Option{
	libp2p.NATPortMap(),
	libp2p.ConnectionManager(connMgr),
	libp2p.EnableAutoRelay(),
	libp2p.EnableNATService(),
}
