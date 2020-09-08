package mpq

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

// static check that Source implements AssetSource
var _ asset.Source = &Source{}

// NewSource creates a new MPQ Source
func NewSource(sourcePath string) (asset.Source, error) {
	loaded, err := d2mpq.Load(sourcePath)
	if err != nil {
		return nil, err
	}

	return &Source{loaded}, nil
}

// Source is an implementation of an asset source for MPQ archives
type Source struct {
	MPQ d2interface.Archive
}

// Type returns the asset type, for MPQ's it always returns the MPQ asset source type
func (v *Source) Type() types.SourceType {
	return types.AssetSourceMPQ
}

// Open attempts to open a file within the MPQ archive
func (v *Source) Open(name string) (a asset.Asset, err error) {
	stream, err := v.MPQ.ReadFileStream(name)

	if err != nil {
		return nil, err
	}

	a = &Asset{
		source: v,
		stream: stream,
	}

	return a, nil
}

// String returns the path of the MPQ on the host filesystem
func (v *Source) String() string {
	return v.MPQ.Path()
}
