# Megasound Documentation

Megasound is a versatile audio processing library written in Go. It provides a rich set of tools for decoding, encoding, manipulating, and playing audio in various formats. This documentation covers all aspects of the Megasound library, including its core interfaces, functionalities, subpackages, and usage examples.

## Table of Contents

- [Megasound Documentation](#megasound-documentation)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Core Package: megasound](#core-package-megasound)
    - [Interfaces](#interfaces)
      - [Streamer](#streamer)
      - [StreamSeeker](#streamseeker)
      - [StreamCloser](#streamcloser)
      - [StreamSeekCloser](#streamseekcloser)
    - [Types](#types)
      - [SampleRate](#samplerate)
      - [Format](#format)
    - [Functions](#functions)
      - [Silence](#silence)
      - [Callback](#callback)
      - [Iterate](#iterate)
      - [Mix](#mix)
      - [Seq](#seq)
      - [Take](#take)
      - [Loop](#loop)
      - [Dup](#dup)
  - [Subpackages](#subpackages)
    - [bpm](#bpm)
        - [Functions](#functions-1)
        - [ProgressivelyReadFloatArray](#progressivelyreadfloatarray)
        - [ScanForBpm](#scanforbpm)
    - [wav](#wav)
      - [Functions](#functions-2)
        - [Encode](#encode)
        - [Decode](#decode)
    - [mp3](#mp3)
      - [Functions](#functions-3)
        - [Decode](#decode-1)
    - [flac](#flac)
      - [Functions](#functions-4)
        - [Decode](#decode-2)
    - [vorbis](#vorbis)
      - [Functions](#functions-5)
        - [Decode](#decode-3)
    - [KeyDetector](#keydetector)
      - [Overview](#overview-1)
        - [Files](#files)
          - [`krumhansl.go`](#krumhanslgo)
        - [Types](#types-1)
        - [KeyResult](#keyresult)
      - [Functions](#functions-6)
        - [`KrumhanslSchmuckler`](#krumhanslschmuckler)
        - [`cosineSimilarity`](#cosinesimilarity)
        - [`camelotWheelMapping`](#camelotwheelmapping)
      - [`pcp.go`](#pcpgo)
        - [Functions](#functions-7)
          - [`ComputePCP`](#computepcp)
        - [`computePitchClass`](#computepitchclass)
        - [`estimateFrequency`](#estimatefrequency)
    - [`keyprofiles.go`](#keyprofilesgo)
      - [Types](#types-2)
        - [KeyProfile](#keyprofile)
      - [Functions](#functions-8)
        - [`PCPKeyProfiles`](#pcpkeyprofiles)
        - [`KrumhanslKeyProfiles`](#krumhanslkeyprofiles)
    - [`keydetector.go`](#keydetectorgo)
      - [Types](#types-3)
        - [KeyDetector](#keydetector-1)
      - [Functions](#functions-9)
        - [`NewKeyDetector`](#newkeydetector)
        - [`DetectKey`](#detectkey)
  - [Effects](#effects)
    - [Overview](#overview-2)
      - [Types](#types-4)
        - [Volume](#volume)
    - [Gain](#gain)
    - [Pan](#pan)
    - [Swap](#swap)
    - [Mono](#mono)
    - [Doppler](#doppler)
    - [Equalizer](#equalizer)
    - [Ctrl](#ctrl)
  - [Functions](#functions-10)
    - [Volume](#volume-1)
    - [Gain](#gain-1)
    - [Pan](#pan-1)
    - [Swap](#swap-1)
    - [Mono](#mono-1)
    - [Doppler](#doppler-1)
    - [Equalizer](#equalizer-1)
    - [Ctrl](#ctrl-1)

## Overview

Megasound is designed to provide a simple yet powerful interface for audio processing in Go. Whether you're building an audio application, a game, or a sound analysis tool, Megasound offers the necessary components to handle audio data seamlessly. It supports multiple audio formats, real-time streaming, audio effects, and playback capabilities.

## Core Package: megasound

### Interfaces

#### Streamer

```go
type Streamer interface {
    Stream(samples [][2]float64) (n int, ok bool)
    Err() error
}
```

**Description:**  
A Streamer is an entity capable of streaming a sequence of audio samples. It can represent both finite and infinite audio sources.

**Methods:**

`Stream(samples [][2]float64) (n int, ok bool)`

- Purpose: Streams up to len(samples) audio samples
- Parameters:
  - samples: A slice of stereo audio samples where each sample is a pair [Left, Right]
- Returns:
  - n: Number of samples streamed
  - ok: Indicates if more samples are available (true) or if the streamer is drained (false)

`Err() error`

- Purpose: Returns any error encountered during streaming
- Returns: An error if one occurred; otherwise, nil

#### StreamSeeker

```go
type StreamSeeker interface {
    Streamer
    Len() int
    Position() int
    Seek(p int) error
}
```

**Description:**  
A StreamSeeker is a finite-duration Streamer that supports seeking to arbitrary positions within the audio stream.

**Additional Methods:**

- `Len() int`: Returns the total number of samples in the stream
- `Position() int`: Returns the current position in the stream (in samples)
- `Seek(p int) error`: Sets the current position to p samples

#### StreamCloser

```go
type StreamCloser interface {
    Streamer
    Close() error
}
```

**Description:**  
A StreamCloser is a Streamer that also manages resources, such as files or network connections, which need to be released when done.

**Additional Method:**

- `Close() error`: Closes the streamer and releases associated resources

#### StreamSeekCloser

```go
type StreamSeekCloser interface {
    Streamer
    Len() int
    Position() int
    Seek(p int) error
    Close() error
}
```

**Description:**  
A StreamSeekCloser combines the functionalities of StreamSeeker and StreamCloser. It represents a finite-duration Streamer that supports seeking and resource management.

### Types

#### SampleRate

```go
type SampleRate int
```

**Description:**  
Represents the number of audio samples per second.

**Methods:**

`D(n int) time.Duration`

- Purpose: Returns the duration of n samples
- Parameters:
  - n: Number of samples
- Returns: A time.Duration representing the time span of n samples

`N(d time.Duration) int`

- Purpose: Calculates the number of samples that span a given duration
- Parameters:
  - d: A time.Duration
- Returns: Number of samples corresponding to duration d

#### Format

```go
type Format struct {
    SampleRate  SampleRate
    NumChannels int
    Precision   int
}
```

**Description:**  
Defines the audio format for a Buffer or any audio source.

**Fields:**

- `SampleRate SampleRate`: Number of samples per second
- `NumChannels int`: Number of audio channels (1 for mono, 2 for stereo, etc.)
- `Precision int`: Number of bytes used to encode a single sample per channel

**Methods:**

`Width() int`

- Description: Returns the number of bytes per frame (all channels)
- Returns: NumChannels * Precision

`EncodeSigned(p []byte, sample [2]float64) int`

- Description: Encodes a stereo sample into signed bytes
- Parameters:
  - p: Byte slice to encode into
  - sample: Stereo sample [Left, Right]
- Returns: Number of bytes written

`EncodeUnsigned(p []byte, sample [2]float64) int`

- Description: Encodes a stereo sample into unsigned bytes
- Parameters:
  - p: Byte slice to encode into
  - sample: Stereo sample [Left, Right]
- Returns: Number of bytes written

`DecodeSigned(p []byte) ([2]float64, int)`

- Description: Decodes signed bytes into a stereo sample
- Parameters:
  - p: Byte slice to decode from
- Returns:
  - sample: Decoded stereo sample [Left, Right]
  - n: Number of bytes read

`DecodeUnsigned(p []byte) ([2]float64, int)`

- Description: Decodes unsigned bytes into a stereo sample
- Parameters:
  - p: Byte slice to decode from
- Returns:
  - sample: Decoded stereo sample [Left, Right]
  - n: Number of bytes read

### Functions

#### Silence

```go
func Silence(num int) Streamer
```

**Description:**  
Creates a Streamer that generates num samples of silence. If num is negative, it generates silence indefinitely.

**Parameters:**

- num: Number of silent samples to generate. Negative for infinite silence

**Returns:**

- A Streamer that outputs silence

#### Callback

```go
func Callback(f func()) Streamer
```

**Description:**  
Creates a Streamer that executes a callback function f the first time its Stream method is called. Subsequent calls will not execute the callback.

**Parameters:**

- f: A function to execute upon the first stream

**Returns:**

- A Streamer that triggers the callback

#### Iterate

```go
func Iterate(g func() Streamer) Streamer
```

**Description:**  
Creates a Streamer that successively streams Streamer instances returned by the generator function g. Streaming stops when g returns nil.

**Parameters:**

- g: A generator function that returns a Streamer

**Returns:**

- A Streamer that streams the sequence of Streamer instances provided by g

#### Mix

```go
func Mix(s ...Streamer) Streamer
```

**Description:**  
Creates a Streamer that mixes multiple Streamer instances together. It does not propagate errors from the individual Streamers.

**Parameters:**

- s: Variadic list of Streamer instances to mix

**Returns:**

- A Streamer that outputs the mixed audio of the provided Streamers

#### Seq

```go
func Seq(s ...Streamer) Streamer
```

**Description:**  
Creates a Streamer that streams multiple Streamer instances in sequence without pauses. It does not propagate errors from the individual Streamers.

**Parameters:**

- s: Variadic list of Streamer instances to sequence

**Returns:**

- A Streamer that outputs the concatenated audio of the provided Streamers

#### Take

```go
func Take(num int, s Streamer) Streamer
```

**Description:**  
Creates a Streamer that streams at most num samples from the source Streamer s.

**Parameters:**

- num: Maximum number of samples to stream
- s: Source Streamer

**Returns:**

- A Streamer that limits the output to num samples from s

#### Loop

```go
func Loop(count int, s StreamSeeker) Streamer
```

**Description:**  
Creates a Streamer that loops the source StreamSeeker s a specified number of times. If count is negative, it loops indefinitely.

**Parameters:**

- count: Number of times to loop. Negative for infinite looping
- s: Source StreamSeeker to loop

**Returns:**

- A Streamer that outputs the looped audio of s

#### Dup

```go
func Dup(s Streamer) (Streamer, Streamer)
```

**Description:**  
Creates two Streamer instances that duplicate the audio from the original Streamer s. The two duplicate Streamers cannot be used concurrently without synchronization.

**Parameters:**

- s: Original Streamer to duplicate

**Returns:**

- Two Streamer instances that output the same audio as s

## Subpackages

### bpm

Package Path: `github.com/rickcollette/megasound/bpm`

**Overview:**  
The bpm package provides functionality for detecting beats per minute (BPM) from audio data. It includes tools for progressive reading of PCM data, energy level calculation, and BPM scanning over specified ranges.

##### Functions

##### ProgressivelyReadFloatArray

```go
func ProgressivelyReadFloatArray(in chan float32, out chan float32)
```

**Description:**  
Processes PCM data into energy levels for BPM analysis in a progressive manner.

**Parameters:**

- `in`: A channel for streaming PCM data as float32 values.
- `out`: A channel for streaming processed energy levels as float32 values.

**Returns:**  
None.

**Usage Example:**

```go
in := make(chan float32)
out := make(chan float32)
go bpm.ProgressivelyReadFloatArray(in, out)
```

##### ScanForBpm

```go
func ScanForBpm(nrg []float32, slowest, fastest float64, steps, samples int) float64
```

**Description:**  
Scans a range of BPM values for the one with the minimum autodifference, based on the energy levels provided.

**Parameters:**

- `nrg`: A slice of energy levels.
- `slowest`: The minimum BPM to scan.
- `fastest`: The maximum BPM to scan.
- `steps`: The number of coarse steps to take in the scanning range.
- `samples`: The number of samples to evaluate per step.

**Returns:**  

- `float64`: The detected BPM.

**Usage Example:**

```go
nrg := []float32{...} // Populate with energy levels
bpm := bpm.ScanForBpm(nrg, 60, 180, 1024, 512)
fmt.Printf("Detected BPM: %.2f
", bpm)
```

### wav

Package Path: `github.com/rickcollette/megasound/wav`

**Overview:**  
The wav package handles encoding and decoding of audio data in the WAVE format. It allows for reading from and writing to WAVE files, supporting various channel configurations and sample precisions.

#### Functions

##### Encode

```go
func Encode(w io.WriteSeeker, s Streamer, format Format) error
```

**Description:**  
Encodes audio data from a Streamer into the WAVE format and writes it to an io.WriteSeeker. The function finalizes the WAVE header after streaming all samples.

**Parameters:**

- w: An io.WriteSeeker where the WAVE data will be written
- s: The source Streamer providing audio samples
- format: The Format specifying the audio format (sample rate, channels, precision)

**Returns:**

- error: Returns an error if encoding fails; otherwise, nil

**Usage Example:**

```go
file, err := os.Create("output.wav")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

streamer := /* your Streamer instance */
format := megasound.Format{
    SampleRate: 44100,
    NumChannels: 2,
    Precision: 2,
}

err = wav.Encode(file, streamer, format)
if err != nil {
    log.Fatal(err)
}
```

##### Decode

```go
func Decode(r io.Reader) (StreamSeekCloser, Format, error)
```

**Description:**  
Decodes WAVE audio data from an io.Reader and returns a StreamSeekCloser for streaming the audio, along with its Format.

**Parameters:**

- r: An io.Reader containing WAVE-encoded audio data

**Returns:**

- StreamSeekCloser: A streamer that allows seeking within the audio
- Format: The audio format details
- error: Returns an error if decoding fails; otherwise, nil

**Usage Example:**

```go
file, err := os.Open("input.wav")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

streamer, format, err := wav.Decode(file)
if err != nil {
    log.Fatal(err)
}
defer streamer.Close()

// Use streamer and format as needed
```

### mp3

Package Path: `github.com/rickcollette/megasound/mp3`

**Overview:**  
The mp3 package facilitates decoding of audio data in the MP3 format. It leverages the go-mp3 library to parse MP3 files and provides a StreamSeekCloser for streaming the decoded audio.

#### Functions

##### Decode

```go
func Decode(rc io.ReadCloser) (StreamSeekCloser, Format, error)
```

**Description:**  
Decodes MP3 audio data from an io.ReadCloser and returns a StreamSeekCloser for streaming the audio, along with its Format.

**Parameters:**

- rc: An io.ReadCloser containing MP3-encoded audio data

**Returns:**

- StreamSeekCloser: A streamer that allows seeking within the audio
- Format: The audio format details
- error: Returns an error if decoding fails; otherwise, nil

**Usage Example:**

```go
file, err := os.Open("input.mp3")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

streamer, format, err := mp3.Decode(file)
if err != nil {
    log.Fatal(err)
}
defer streamer.Close()

// Use streamer and format as needed
```

### flac

Package Path: `github.com/rickcollette/megasound/flac`

**Overview:**  
The flac package handles decoding of audio data in the FLAC format. It utilizes the mewkiz/flac library to parse FLAC files and provides a StreamSeekCloser for streaming the decoded audio.

#### Functions

##### Decode

```go
func Decode(r io.Reader) (StreamSeekCloser, Format, error)
```

**Description:**  
Decodes FLAC audio data from an io.Reader and returns a StreamSeekCloser for streaming the audio, along with its Format.

**Parameters:**

- r: An io.Reader containing FLAC-encoded audio data

**Returns:**

- StreamSeekCloser: A streamer that allows seeking within the audio
- Format: The audio format details
- error: Returns an error if decoding fails; otherwise, nil

**Usage Example:**

```go
file, err := os.Open("input.flac")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

streamer, format, err := flac.Decode(file)
if err != nil {
    log.Fatal(err)
}
defer streamer.Close()

// Use streamer and format as needed
```

### vorbis

Package Path: `github.com/rickcollette/megasound/vorbis`

**Overview:**  
The vorbis package manages decoding of audio data in the Ogg Vorbis format. It leverages the jfreymuth/oggvorbis library to parse Vorbis streams and provides a StreamSeekCloser for streaming the decoded audio.

#### Functions

##### Decode

```go
func Decode(r io.Reader) (StreamSeekCloser, Format, error)
```

**Description:**  
Decodes Ogg Vorbis audio data from an io.Reader and returns a StreamSeekCloser for streaming the audio, along with its Format.

**Parameters:**

- r: An io.Reader containing Ogg Vorbis-encoded audio data

**Returns:**

- StreamSeekCloser: A streamer that allows seeking within the audio
- Format: The audio format details
- error: Returns an error if decoding fails; otherwise, nil

**Usage Example:**

```go
file, err := os.Open("input.ogg")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

streamer, format, err := vorbis.Decode(file)
if err != nil {
    log.Fatal(err)
}
defer streamer.Close()

// Use streamer and format as needed
```

### KeyDetector

#### Overview

The `keydetector` package provides tools for determining the musical key of an audio track based on its Pitch Class Profile (PCP). It includes utilities for PCP computation, similarity calculation, and key detection with Camelot Wheel notation.

---

##### Files

###### `krumhansl.go`

**Overview:**  
Defines the main functions for detecting the musical key using the Krumhansl-Schmuckler algorithm.

##### Types

##### KeyResult

```go
type KeyResult struct {
    Key             string
    CamelotNotation string
    Confidence      float64
}
```

**Description:**  
Represents the detected key along with its Camelot Wheel notation and a confidence score.

#### Functions

##### `KrumhanslSchmuckler`

```go
func KrumhanslSchmuckler(pcp []float64) KeyResult
```

**Description:**  
Computes the most likely musical key based on the provided PCP vector.

**Parameters:**

- `pcp`: A slice of 12 floats representing the Pitch Class Profile.

**Returns:**  
A `KeyResult` containing the detected key, its Camelot Wheel notation, and confidence score.

**Usage Example:**

```go
pcp := []float64{...} // Populate with PCP values
result := KrumhanslSchmuckler(pcp)
fmt.Printf("Key: %s, Camelot: %s, Confidence: %.2f\n", result.Key, result.CamelotNotation, result.Confidence)
```

##### `cosineSimilarity`

```go
func cosineSimilarity(a, b []float64) float64
```

**Description:**  
Calculates the cosine similarity between two vectors.

**Parameters:**

- `a`: First vector of floats.
- `b`: Second vector of floats.

**Returns:**  
A `float64` value representing the similarity.

**Usage Example:**

```go
similarity := cosineSimilarity(vectorA, vectorB)
fmt.Printf("Cosine Similarity: %.2f\n", similarity)
```

##### `camelotWheelMapping`

```go
func camelotWheelMapping() map[string]string
```

**Description:**  
Returns a map of musical keys to their Camelot Wheel notation.

---

#### `pcp.go`

**Overview:**  
Handles the computation of the Pitch Class Profile (PCP) from audio streams.

##### Functions

###### `ComputePCP`

```go
func ComputePCP(streamer megasound.Streamer, sampleRate int) ([]float64, error)
```

**Description:**  
Computes the PCP for an audio stream.

**Parameters:**

- `streamer`: An audio streamer from which audio samples are retrieved.
- `sampleRate`: The sample rate of the audio.

**Returns:**  
A slice of PCP values and an error if computation fails.

**Usage Example:**

```go
pcp, err := ComputePCP(streamer, 44100)
if err != nil {
    log.Fatalf("Error computing PCP: %v", err)
}
fmt.Printf("PCP: %v\n", pcp)
```

##### `computePitchClass`

```go
func computePitchClass(frequency float64) int
```

**Description:**  
Calculates the pitch class (0-11) for a given frequency.

**Parameters:**

- `frequency`: The frequency of the audio in Hz.

**Returns:**  
An integer representing the pitch class.

##### `estimateFrequency`

```go
func estimateFrequency(samples [][2]float64, sampleRate int, mono []float64) float64
```

**Description:**  
Estimates the fundamental frequency from audio samples using FFT.

**Parameters:**

- `samples`: Stereo audio samples.
- `sampleRate`: The sample rate of the audio.
- `mono`: Preallocated buffer for mono samples.

**Returns:**  
The estimated frequency in Hz.

---

### `keyprofiles.go`

**Overview:**  
Provides predefined key profiles for major and minor keys based on music theory.

#### Types

##### KeyProfile

```go
type KeyProfile struct {
    Profile []float64
    Norm    float64
}
```

**Description:**  
Represents a key profile with its precomputed norm.

#### Functions

##### `PCPKeyProfiles`

```go
func PCPKeyProfiles() map[string][]float64
```

**Description:**  
Returns predefined key profiles for all major and minor keys.

**Returns:**  
A map where keys are string names of musical keys and values are slices of floats representing the profiles.

**Usage Example:**

```go
profiles := PCPKeyProfiles()
fmt.Printf("C Major Profile: %v\n", profiles["C Major"])
```

##### `KrumhanslKeyProfiles`

```go
func KrumhanslKeyProfiles() map[string][]float64
```

**Description:**  
Provides key profiles based on the Krumhansl-Schmuckler algorithm.

**Returns:**  
A map of musical keys to their profiles.

**Usage Example:**

```go
profiles := KrumhanslKeyProfiles()
fmt.Printf("A Minor Profile: %v\n", profiles["A Minor"])
```

---

### `keydetector.go`

**Overview:**  
The entry point for key detection, combining PCP computation and key analysis.

#### Types

##### KeyDetector

```go
type KeyDetector struct {
    Streamer   megasound.Streamer
    SampleRate int
}
```

**Description:**  
Encapsulates the functionality for detecting musical keys.

#### Functions

##### `NewKeyDetector`

```go
func NewKeyDetector(s megasound.Streamer, sampleRate int) *KeyDetector
```

**Description:**  
Creates a new instance of `KeyDetector`.

##### `DetectKey`

```go
func (kd *KeyDetector) DetectKey() (KeyResult, error)
```

**Description:**  
Detects the musical key from the audio stream.

**Returns:**  
A `KeyResult` with the detected key, Camelot Wheel notation, and confidence.

**Usage Example:**

```go
kd := NewKeyDetector(streamer, 44100)
result, err := kd.DetectKey()
if err != nil {
    log.Fatalf("Error detecting key: %v", err)
}
fmt.Printf("Key: %s, Camelot: %s, Confidence: %.2f\n", result.Key, result.CamelotNotation, result.Confidence)
```

## Effects

**Package Path:** `github.com/rickcollette/megasound/effects`

### Overview

The `effects` package provides a collection of audio effects that can be applied to `Streamer` instances. These effects include volume control, panning, mono conversion, gain adjustment, Doppler effect simulation, and more.

---

#### Types

##### Volume

```go
type Volume struct {
    Streamer Streamer
    Base     float64
    Volume   float64
    Silent   bool
}
```

**Description:**  
Adjusts the volume of the wrapped `Streamer` in a human-natural way using exponential gain. When `Silent` is set to `true`, the output is muted.

**Fields:**

- `Streamer`: The source `Streamer` to adjust.
- `Base`: The base used for exponential gain calculation.
- `Volume`: The exponent applied to the base.
- `Silent`: If `true`, the output is muted.

---

### Gain

```go
type Gain struct {
    Streamer Streamer
    Gain     float64
}
```

**Description:**  
Applies linear gain to the wrapped `Streamer`. The output is multiplied by `1 + Gain`.

**Fields:**

- `Streamer`: The source `Streamer` to amplify.
- `Gain`: The linear gain value.

---

### Pan

```go
type Pan struct {
    Streamer Streamer
    Pan      float64
}
```

**Description:**  
Balances the audio between the left and right channels. A `Pan` value of `-1` sends all audio to the left channel, `+1` to the right channel, and `0` maintains the original balance.

**Fields:**

- `Streamer`: The source `Streamer` to pan.
- `Pan`: Panning value ranging from `-1` (left) to `+1` (right).

---

### Swap

```go
type swap struct {
    Streamer Streamer
}
```

**Description:**  
Swaps the left and right audio channels of the wrapped `Streamer`.

**Fields:**

- `Streamer`: The source `Streamer` to swap channels.

---

### Mono

```go
type mono struct {
    Streamer Streamer
}
```

**Description:**  
Converts stereo audio to mono by averaging the left and right channels.

**Fields:**

- `Streamer`: The source `Streamer` to convert to mono.

---

### Doppler

```go
type doppler struct {
    r               *Resampler
    distance        func(delta int) float64
    space           [][2]float64
    samplesPerMeter float64
}
```

**Description:**  
Simulates the Doppler effect, adjusting the density of the audio stream based on the distance from the sound source.

**Fields:**

- `r`: Underlying `Resampler`.
- `distance`: Function to calculate the current distance based on the number of samples.
- `space`: Buffer for resampled audio data.
- `samplesPerMeter`: Ratio of sample rate to the speed of sound.

---

### Equalizer

```go
type equalizer struct {
    streamer Streamer
    sections []section
}
```

**Description:**  
Applies a parametric equalizer to the audio stream, allowing for frequency-based adjustments.

**Fields:**

- `streamer`: The source `Streamer` to equalize.
- `sections`: Equalizer sections defining frequency adjustments.

---

### Ctrl

```go
type Ctrl struct {
    Streamer Streamer
    Paused   bool
}
```

**Description:**  
Allows pausing and resuming of a `Streamer`. When paused, the `Ctrl` streams silence.

**Fields:**

- `Streamer`: The source `Streamer` to control.
- `Paused`: If `true`, the `Streamer` is paused and outputs silence.

---

## Functions

### Volume

```go
func Volume(s Streamer, base, volume float64, silent bool) *Volume
```

**Description:**  
Creates a `Volume` effect to adjust the audio volume exponentially.

**Parameters:**

- `s`: Source `Streamer`.
- `base`: Base for exponential calculation (commonly 2).
- `volume`: Exponent value (e.g., `dB/10` for decibel adjustments).
- `silent`: Mutes the output if `true`.

**Returns:**  
A `Volume` instance to apply the effect.

---

### Gain

```go
func Gain(s Streamer, gain float64) *Gain
```

**Description:**  
Creates a `Gain` effect to apply linear amplification to the audio.

**Parameters:**

- `s`: Source `Streamer`.
- `gain`: Gain value to apply.

**Returns:**  
A `Gain` instance to apply the effect.

---

### Pan

```go
func Pan(s Streamer, pan float64) *Pan
```

**Description:**  
Creates a `Pan` effect to balance audio between left and right channels.

**Parameters:**

- `s`: Source `Streamer`.
- `pan`: Panning value ranging from `-1` (left) to `+1` (right).

**Returns:**  
A `Pan` instance to apply the effect.

---

### Swap

```go
func Swap(s Streamer) Streamer
```

**Description:**  
Swaps the left and right audio channels of the provided `Streamer`.

**Parameters:**

- `s`: Source `Streamer`.

**Returns:**  
A `Streamer` with swapped channels.

---

### Mono

```go
func Mono(s Streamer) Streamer
```

**Description:**  
Converts stereo audio to mono by averaging the left and right channels.

**Parameters:**

- `s`: Source `Streamer`.

**Returns:**  
A `Streamer` that outputs mono audio.

---

### Doppler

```go
func Doppler(quality int, samplesPerMeter float64, s Streamer, distance func(delta int) float64) Streamer
```

**Description:**  
Simulates the Doppler effect by adjusting the audio density based on dynamic distance changes.

**Parameters:**

- `quality`: Quality of the underlying resampler (1 or 2 recommended).
- `samplesPerMeter`: Ratio of sample rate to the speed of sound.
- `s`: Source `Streamer`.
- `distance`: Function to calculate current distance based on the number of samples.

**Returns:**  
A `Streamer` that applies the Doppler effect.

---

### Equalizer

```go
func NewEqualizer(s Streamer, sr SampleRate, sections EqualizerSections) Streamer
```

**Description:**  
Creates an Equalizer effect to apply frequency-based adjustments to the audio stream.

**Parameters:**

- `s`: Source `Streamer`.
- `sr`: Sample rate of the audio stream.
- `sections`: Equalizer sections defining frequency adjustments.

**Returns:**  
A `Streamer` with equalization applied.

---

### Ctrl

```go
func (c *Ctrl) Stream(samples [][2]float64) (n int, ok bool)
```

**Description:**  
Streams audio samples from the controlled `Streamer`. If paused, streams silence.

**Parameters:**

- `samples`: Slice to populate with audio samples.

**Returns:**

- `n`: Number of samples streamed.
- `ok`: Indicates if more samples are available.

---
