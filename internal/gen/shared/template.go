// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package shared

type BlendTemplate struct {
	Name         string
	Kernel       string
	EquationRGBA string
	KernelDocs   string
}

// BlendTemplates is the list of all the blend modes that can be generated.
//
// Each template is composed of a kernel and an optional equation.
// The kernel is used for non-premultiplied logic (NRGBA or converted RGBA).
// The equation is used for a fast path when the premultiplied logic is used (RGBA only).
//
// Both the kernel and equation are code snippets for color-channel calculations with specific placeholders.
// The placeholders are repeated in the output for each color channel (RGB).
//
// Kernel placeholders:
// $R: the result/output color channel
// $S: the source color channel
// $D: the destination color channel
//
// Equation placeholders:
// $Sp: the premultiplied source color channel
// $Dp: the premultiplied destination color channel
// $sA: the source alpha channel
// $dA: the destination alpha channel
//
// Available optimized custom helper functions:
// md255: multiply two values and divide by 255; input must be [0, 255]
// sqrt: square root of a color channel; input must be [0, 255]
// div255: divide value by 255 (approximated)
// unpremultiply: un-premultiply a color channel
var BlendTemplates = []BlendTemplate{
	{
		Name:       "ColorBurn",
		Kernel:     "if $S + $D <= 255 { $R = 0 } else { $R = 255 - ((255 - $D) * 255 + $S/2) / $S }",
		KernelDocs: "Cr = 1 - (1 - Cd) / Cs",
	},
	{
		Name:       "ColorDodge",
		Kernel:     "if $S == 255 { $R = 255 } else { $R = min(($D * 255) / (255 - $S), 255) }",
		KernelDocs: "Cr = Cd / (1 - Cs)",
	},
	{
		Name:       "Darken",
		Kernel:     "$R = min($S, $D)",
		KernelDocs: "Cr = min(Cs, Cd)",
	},
	{
		Name:         "Difference",
		Kernel:       "$R = min($S - $D, $D - $S)",
		KernelDocs:   "Cr = abs(Cs - Cd)",
		EquationRGBA: "$Sp + $Dp - div255(2*min($Sp*$dA, $Dp*$sA))",
	},
	{
		Name:       "Divide",
		Kernel:     "if $S == 0 { $R = 255 } else { $R = min(($D*255 + $S/2)/$S, 255) }",
		KernelDocs: "Cr = Cd / Cs",
	},
	{
		Name:         "Exclusion",
		Kernel:       "$R = $S + $D - 2*md255($S, $D)",
		KernelDocs:   "Cr = Cs + Cd - 2 * Cs * Cd",
		EquationRGBA: "$Sp + $Dp - 2*md255($Sp, $Dp)",
	},
	{
		Name:       "HardLight",
		Kernel:     "if $S < 128 { $R = 2 * md255($S, $D) } else { $R = 255 - 2*md255(255-$S, 255-$D) }",
		KernelDocs: "if Cs < 0.5 { Cr = 2 * Cs * Cd } else { Cr = 1 - 2 * (1 - Cs) * (1 - Cd) }",
	},
	{
		Name:       "HardMix",
		Kernel:     "if $S + $D < 255 { $R = 0 } else { $R = 255 }",
		KernelDocs: "if Cs + Cd < 1 { Cr = 0 } else { Cr = 1 }",
	},
	{
		Name:       "Lighten",
		Kernel:     "$R = max($S, $D)",
		KernelDocs: "Cr = max(Cs, Cd)",
	},
	{
		Name:       "LinearBurn",
		Kernel:     "$R = $S + $D - 255; if $R > 255 { $R = 0 }",
		KernelDocs: "Cr = Cs + Cd - 1",
	},
	{
		Name:       "LinearDodge",
		Kernel:     "$R = min($S + $D, 255)",
		KernelDocs: "Cr = Cs + Cd",
	},
	{
		Name:       "LinearLight",
		Kernel:     "$R = $D + 2*$S - 255; if $R > 510 { $R = 0 } else if $R > 255 { $R = 255 }",
		KernelDocs: "Cr = Cd + 2*Cs - 1",
	},
	{
		Name:         "Multiply",
		Kernel:       "$R = md255($S, $D)",
		KernelDocs:   "Cr = Cs * Cd",
		EquationRGBA: "md255($Dp, 255-$sA) + md255($Sp, 255-$dA) + md255($Sp, $Dp)",
	},
	{
		Name:       "Overlay",
		Kernel:     "if $D < 128 { $R = 2 * md255($S, $D) } else { $R = 255 - 2*md255(255-$S, 255-$D) }",
		KernelDocs: "if Cd < 0.5 { Cr = 2 * Cs * Cd } else { Cr = 1 - 2 * (1 - Cs) * (1 - Cd) }",
	},
	{
		Name:       "PinLight",
		Kernel:     "if $S < 128 { $R = min($D, 2 * $S) } else { $R = max($D, 2*$S-255) }",
		KernelDocs: "if Cs < 0.5 { Cr = min(Cd, 2 * Cs) } else { Cr = max(Cd, 2 * Cs - 1) }",
	},
	{
		Name:         "Screen",
		Kernel:       "$R = 255 - md255(255-$S, 255-$D)",
		KernelDocs:   "Cr = 1 - (1 - Cs) * (1 - Cd)",
		EquationRGBA: "$Sp + $Dp - md255($Sp, $Dp)",
	},
	{
		Name:       "SoftLight",
		Kernel:     "if $S < 128 { $R = $D - div255(div255((255 - 2*$S) * $D * (255 - $D))) } else { $R = $D + md255(2*$S - 255, sqrt($D) - $D) }",
		KernelDocs: "if Cs < 0.5 { Cr = Cd - (1 - 2*Cs) * Cd * (1 - Cd) } else { Cr = Cd + (2*Cs - 1) * (sqrt(Cd) - Cd) }",
	},
	{
		Name:       "Subtract",
		Kernel:     "$R = uint32(max(int($D) - int($S), 0))",
		KernelDocs: "Cr = Cd - Cs",
	},
	{
		Name: "VividLight",
		Kernel: "switch { " +
			"case $S == 0 || $S == 255: $R = $S; " +
			"case $S < 128: $R = 255 - min(((255 - $D) * 255 + $S) / (2 * $S), 255); " +
			"default: $R = min((($D * 255) + (255 - $S)) / (510 - 2 * $S), 255) }",
		KernelDocs: "if Cs < 0.5 { Cr = 1 - (1 - Cd) / (2 * Cs) } else { Cr = Cd / (2 * (1 - Cs)) }",
	},
}
