// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eth

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/params"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		NoPruning               bool
		NoPrefetch              bool
		Whitelist               map[uint64]common.Hash `toml:"-"`
		LightServ               int                    `toml:",omitempty"`
		LightIngress            int                    `toml:",omitempty"`
		LightEgress             int                    `toml:",omitempty"`
		LightPeers              int                    `toml:",omitempty"`
		GatewayFee              *big.Int               `toml:",omitempty"`
		Etherbase               common.Address         `toml:",omitempty"`
		BLSbase                 common.Address         `toml:",omitempty"`
		UltraLightServers       []string               `toml:",omitempty"`
		UltraLightFraction      int                    `toml:",omitempty"`
		UltraLightOnlyAnnounce  bool                   `toml:",omitempty"`
		SkipBcVersionCheck      bool                   `toml:"-"`
		DatabaseHandles         int                    `toml:"-"`
		DatabaseCache           int
		DatabaseFreezer         string
		TrieCleanCache          int
		TrieDirtyCache          int
		TrieTimeout             time.Duration
		Miner                   miner.Config
		TxPool                  core.TxPoolConfig
		EnablePreimageRecording bool
		Istanbul                istanbul.Config
		DocRoot                 string `toml:"-"`
		EWASMInterpreter        string
		EVMInterpreter          string
		RPCGasCap               *big.Int                       `toml:",omitempty"`
		Checkpoint              *params.TrustedCheckpoint      `toml:",omitempty"`
		CheckpointOracle        *params.CheckpointOracleConfig `toml:",omitempty"`
		OverrideIstanbul        *big.Int `toml:",omitempty"`

	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.NoPruning = c.NoPruning
	enc.NoPrefetch = c.NoPrefetch
	enc.Whitelist = c.Whitelist
	enc.LightServ = c.LightServ
	enc.LightIngress = c.LightIngress
	enc.LightEgress = c.LightEgress
	enc.LightPeers = c.LightPeers
	enc.GatewayFee = c.GatewayFee
	enc.Etherbase = c.Etherbase
	enc.BLSbase = c.BLSbase
	enc.UltraLightServers = c.UltraLightServers
	enc.UltraLightFraction = c.UltraLightFraction
	enc.UltraLightOnlyAnnounce = c.UltraLightOnlyAnnounce
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.DatabaseFreezer = c.DatabaseFreezer
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.Miner = c.Miner
	enc.TxPool = c.TxPool
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.Istanbul = c.Istanbul
	enc.DocRoot = c.DocRoot
	enc.EWASMInterpreter = c.EWASMInterpreter
	enc.EVMInterpreter = c.EVMInterpreter
	enc.RPCGasCap = c.RPCGasCap
	enc.Checkpoint = c.Checkpoint
	enc.CheckpointOracle = c.CheckpointOracle
	enc.OverrideIstanbul = c.OverrideIstanbul
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		NoPruning               *bool
		NoPrefetch              *bool
		Whitelist               map[uint64]common.Hash `toml:"-"`
		LightServ               *int                   `toml:",omitempty"`
		LightIngress            *int                   `toml:",omitempty"`
		LightEgress             *int                   `toml:",omitempty"`
		LightPeers              *int                   `toml:",omitempty"`
		GatewayFee              *big.Int               `toml:",omitempty"`
		Etherbase               *common.Address        `toml:",omitempty"`
		BLSbase                 *common.Address        `toml:",omitempty"`
		UltraLightServers       []string               `toml:",omitempty"`
		UltraLightFraction      *int                   `toml:",omitempty"`
		UltraLightOnlyAnnounce  *bool                  `toml:",omitempty"`
		SkipBcVersionCheck      *bool                  `toml:"-"`
		DatabaseHandles         *int                   `toml:"-"`
		DatabaseCache           *int
		DatabaseFreezer         *string
		TrieCleanCache          *int
		TrieDirtyCache          *int
		TrieTimeout             *time.Duration
		Miner                   *miner.Config
		TxPool                  *core.TxPoolConfig
		EnablePreimageRecording *bool
		Istanbul                *istanbul.Config
		DocRoot                 *string `toml:"-"`
		EWASMInterpreter        *string
		EVMInterpreter          *string
		RPCGasCap               *big.Int                       `toml:",omitempty"`
		Checkpoint              *params.TrustedCheckpoint      `toml:",omitempty"`
		CheckpointOracle        *params.CheckpointOracleConfig `toml:",omitempty"`
		OverrideIstanbul        *big.Int `toml:",omitempty"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.NoPrefetch != nil {
		c.NoPrefetch = *dec.NoPrefetch
	}
	if dec.Whitelist != nil {
		c.Whitelist = dec.Whitelist
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightIngress != nil {
		c.LightIngress = *dec.LightIngress
	}
	if dec.LightEgress != nil {
		c.LightEgress = *dec.LightEgress
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.GatewayFee != nil {
		c.GatewayFee = dec.GatewayFee
	}
	if dec.Etherbase != nil {
		c.Etherbase = *dec.Etherbase
	}
	if dec.BLSbase != nil {
		c.BLSbase = *dec.BLSbase
	}
	if dec.UltraLightServers != nil {
		c.UltraLightServers = dec.UltraLightServers
	}
	if dec.UltraLightFraction != nil {
		c.UltraLightFraction = *dec.UltraLightFraction
	}
	if dec.UltraLightOnlyAnnounce != nil {
		c.UltraLightOnlyAnnounce = *dec.UltraLightOnlyAnnounce
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.DatabaseFreezer != nil {
		c.DatabaseFreezer = *dec.DatabaseFreezer
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.Istanbul != nil {
		c.Istanbul = *dec.Istanbul
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.EWASMInterpreter != nil {
		c.EWASMInterpreter = *dec.EWASMInterpreter
	}
	if dec.EVMInterpreter != nil {
		c.EVMInterpreter = *dec.EVMInterpreter
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = dec.RPCGasCap
	}
	if dec.Checkpoint != nil {
		c.Checkpoint = dec.Checkpoint
	}
	if dec.CheckpointOracle != nil {
		c.CheckpointOracle = dec.CheckpointOracle
	}
	if dec.OverrideIstanbul != nil {
		c.OverrideIstanbul = dec.OverrideIstanbul
	}
	return nil
}
