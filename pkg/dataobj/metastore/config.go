package metastore

import (
	"flag"
	"time"
)

// Config is the configuration block for the metastore settings.
type Config struct {
	IndexStoragePrefix string `yaml:"index_storage_prefix" experimental:"true"`

	TOCResolver TOCResolverConfig `yaml:"toc_resolver" category:"experimental"`
}

// RegisterFlags registers the flags for the metastore settings.
func (c *Config) RegisterFlags(f *flag.FlagSet) {
	prefix := "dataobj-metastore."
	f.StringVar(&c.IndexStoragePrefix, prefix+"index-storage-prefix", "index/v0", "Experimental: A prefix to use for storing indexes in object storage. Used for testing only.")

	c.TOCResolver.RegisterFlagsWithPrefix(prefix+"toc-resolver.", f)
}

// Validate returns an error when the settings are invalid. storageLag is the delay after which data objects hold all the
// data.
func (c *Config) Validate(storageLag time.Duration) error {
	return c.TOCResolver.Validate(storageLag)
}
