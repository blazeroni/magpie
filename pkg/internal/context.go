// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package internal

import (
	"fmt"
	"image/color"
	"runtime"

	"github.com/blazeroni/magpie/pkg/core"
)

var _ core.Config = (*Config)(nil)

// DefaultConfig is the configuration used by top-level functions.
var DefaultConfig *Config

// Configures the default configuration.
// Can be overridden by passing a new Config to SetDefaultConfig.
func init() {
	// default to half the available max
	goroutines := max(1, runtime.GOMAXPROCS(0)/2)
	config := Config{
		pixelIterator:     core.NewPixelIterator(goroutines),
		defaultOutputMode: core.DefaultOutputToDst,
		defaultColorModel: color.NRGBAModel,
	}
	DefaultConfig = &config
}

type Config struct {
	pixelIterator     core.PixelIterator
	defaultColorModel color.Model
	defaultOutputMode core.DefaultOutputMode
}

func NewConfig(pixIter core.PixelIterator, defaultOutputMode core.DefaultOutputMode, defaultColorModel color.Model) *Config {
	return &Config{
		pixelIterator:     pixIter,
		defaultOutputMode: defaultOutputMode,
		defaultColorModel: defaultColorModel,
	}
}

func (c *Config) PixelIterator() core.PixelIterator {
	return c.pixelIterator
}

func (c *Config) DefaultOutputMode() core.DefaultOutputMode {
	return c.defaultOutputMode
}

func (c *Config) SetPixelIterator(pixIter core.PixelIterator) {
	c.pixelIterator = pixIter
}

func (c *Config) SetDefaultOutputMode(defaultOutput core.DefaultOutputMode) {
	c.defaultOutputMode = defaultOutput
}

func (c *Config) SetDefaultColorModel(model color.Model) error {
	if !core.IsColorModelSupported(model) {
		return fmt.Errorf("unsupported color model")
	}
	c.defaultColorModel = model
	return nil
}

func (c *Config) DefaultColorModel() color.Model {
	return c.defaultColorModel
}
