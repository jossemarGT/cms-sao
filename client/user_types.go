// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": Application User Types
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.4.1

package client

import (
	"github.com/goadesign/goa"
)

// Abstracts the common attributes from EntryPayload and EntryFormPayload
type abstractEntry struct {
	// Contest unique and human readable string identifier
	ContestSlug *string `form:"contestSlug,omitempty" json:"contestSlug,omitempty" yaml:"contestSlug,omitempty" xml:"contestSlug,omitempty"`
	// Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words
	// 		any submitted Entry will use a CMS Token
	Ranked *bool `form:"ranked,omitempty" json:"ranked,omitempty" yaml:"ranked,omitempty" xml:"ranked,omitempty"`
	// Task unique and human readable string identifier
	TaskSlug *string `form:"taskSlug,omitempty" json:"taskSlug,omitempty" yaml:"taskSlug,omitempty" xml:"taskSlug,omitempty"`
}

// Finalize sets the default values for abstractEntry type instance.
func (ut *abstractEntry) Finalize() {
	var defaultContestSlug = ""
	if ut.ContestSlug == nil {
		ut.ContestSlug = &defaultContestSlug
	}
	var defaultRanked = true
	if ut.Ranked == nil {
		ut.Ranked = &defaultRanked
	}
	var defaultTaskSlug = ""
	if ut.TaskSlug == nil {
		ut.TaskSlug = &defaultTaskSlug
	}
}

// Validate validates the abstractEntry type instance.
func (ut *abstractEntry) Validate() (err error) {
	if ut.ContestSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *ut.ContestSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.contestSlug`, *ut.ContestSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	if ut.TaskSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *ut.TaskSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.taskSlug`, *ut.TaskSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	return
}

// Publicize creates AbstractEntry from abstractEntry
func (ut *abstractEntry) Publicize() *AbstractEntry {
	var pub AbstractEntry
	if ut.ContestSlug != nil {
		pub.ContestSlug = *ut.ContestSlug
	}
	if ut.Ranked != nil {
		pub.Ranked = *ut.Ranked
	}
	if ut.TaskSlug != nil {
		pub.TaskSlug = *ut.TaskSlug
	}
	return &pub
}

// Abstracts the common attributes from EntryPayload and EntryFormPayload
type AbstractEntry struct {
	// Contest unique and human readable string identifier
	ContestSlug string `form:"contestSlug" json:"contestSlug" yaml:"contestSlug" xml:"contestSlug"`
	// Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words
	// 		any submitted Entry will use a CMS Token
	Ranked bool `form:"ranked" json:"ranked" yaml:"ranked" xml:"ranked"`
	// Task unique and human readable string identifier
	TaskSlug string `form:"taskSlug" json:"taskSlug" yaml:"taskSlug" xml:"taskSlug"`
}

// Validate validates the AbstractEntry type instance.
func (ut *AbstractEntry) Validate() (err error) {
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, ut.ContestSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.contestSlug`, ut.ContestSlug, `[_a-zA-Z0-9\-]+`))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, ut.TaskSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.taskSlug`, ut.TaskSlug, `[_a-zA-Z0-9\-]+`))
	}
	return
}

// Embedded reprensentation of an entry compilation result
type compilationResult struct {
	// Memory consumed
	Memory *int `form:"memory,omitempty" json:"memory,omitempty" yaml:"memory,omitempty" xml:"memory,omitempty"`
	// Execution result status
	Status *string `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
	// Compilation process' standard error
	Stderr *string `form:"stderr,omitempty" json:"stderr,omitempty" yaml:"stderr,omitempty" xml:"stderr,omitempty"`
	// Compilation process' standard output
	Stdout *string `form:"stdout,omitempty" json:"stdout,omitempty" yaml:"stdout,omitempty" xml:"stdout,omitempty"`
	// The spent execution CPU time
	Time *float64 `form:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty" xml:"time,omitempty"`
	// Compilation retries
	Tries *int `form:"tries,omitempty" json:"tries,omitempty" yaml:"tries,omitempty" xml:"tries,omitempty"`
	// The spent execution human perceived time
	WallClockTime *float64 `form:"wallClockTime,omitempty" json:"wallClockTime,omitempty" yaml:"wallClockTime,omitempty" xml:"wallClockTime,omitempty"`
}

// Finalize sets the default values for compilationResult type instance.
func (ut *compilationResult) Finalize() {
	var defaultStatus = "unprocessed"
	if ut.Status == nil {
		ut.Status = &defaultStatus
	}
}

// Validate validates the compilationResult type instance.
func (ut *compilationResult) Validate() (err error) {
	if ut.Status != nil {
		if !(*ut.Status == "ok" || *ut.Status == "fail" || *ut.Status == "unprocessed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.status`, *ut.Status, []interface{}{"ok", "fail", "unprocessed"}))
		}
	}
	if ut.Tries != nil {
		if *ut.Tries < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.tries`, *ut.Tries, 0, true))
		}
	}
	return
}

// Publicize creates CompilationResult from compilationResult
func (ut *compilationResult) Publicize() *CompilationResult {
	var pub CompilationResult
	if ut.Memory != nil {
		pub.Memory = ut.Memory
	}
	if ut.Status != nil {
		pub.Status = *ut.Status
	}
	if ut.Stderr != nil {
		pub.Stderr = ut.Stderr
	}
	if ut.Stdout != nil {
		pub.Stdout = ut.Stdout
	}
	if ut.Time != nil {
		pub.Time = ut.Time
	}
	if ut.Tries != nil {
		pub.Tries = ut.Tries
	}
	if ut.WallClockTime != nil {
		pub.WallClockTime = ut.WallClockTime
	}
	return &pub
}

// Embedded reprensentation of an entry compilation result
type CompilationResult struct {
	// Memory consumed
	Memory *int `form:"memory,omitempty" json:"memory,omitempty" yaml:"memory,omitempty" xml:"memory,omitempty"`
	// Execution result status
	Status string `form:"status" json:"status" yaml:"status" xml:"status"`
	// Compilation process' standard error
	Stderr *string `form:"stderr,omitempty" json:"stderr,omitempty" yaml:"stderr,omitempty" xml:"stderr,omitempty"`
	// Compilation process' standard output
	Stdout *string `form:"stdout,omitempty" json:"stdout,omitempty" yaml:"stdout,omitempty" xml:"stdout,omitempty"`
	// The spent execution CPU time
	Time *float64 `form:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty" xml:"time,omitempty"`
	// Compilation retries
	Tries *int `form:"tries,omitempty" json:"tries,omitempty" yaml:"tries,omitempty" xml:"tries,omitempty"`
	// The spent execution human perceived time
	WallClockTime *float64 `form:"wallClockTime,omitempty" json:"wallClockTime,omitempty" yaml:"wallClockTime,omitempty" xml:"wallClockTime,omitempty"`
}

// Validate validates the CompilationResult type instance.
func (ut *CompilationResult) Validate() (err error) {
	if !(ut.Status == "ok" || ut.Status == "fail" || ut.Status == "unprocessed") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.status`, ut.Status, []interface{}{"ok", "fail", "unprocessed"}))
	}
	if ut.Tries != nil {
		if *ut.Tries < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.tries`, *ut.Tries, 0, true))
		}
	}
	return
}

// Any source code or input that should be compiled, executed or evaluated
type entryFormPayload struct {
	// Contest unique and human readable string identifier
	ContestSlug *string `form:"contestSlug,omitempty" json:"contestSlug,omitempty" yaml:"contestSlug,omitempty" xml:"contestSlug,omitempty"`
	// Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words
	// 		any submitted Entry will use a CMS Token
	Ranked *bool `form:"ranked,omitempty" json:"ranked,omitempty" yaml:"ranked,omitempty" xml:"ranked,omitempty"`
	// Source files representation. Within this list the source code files and input files can be
	// 						  sent alike.
	Sources []multipart.FileHeader `form:"sources,omitempty" json:"sources,omitempty" yaml:"sources,omitempty" xml:"sources,omitempty"`
	// Task unique and human readable string identifier
	TaskSlug *string `form:"taskSlug,omitempty" json:"taskSlug,omitempty" yaml:"taskSlug,omitempty" xml:"taskSlug,omitempty"`
}

// Finalize sets the default values for entryFormPayload type instance.
func (ut *entryFormPayload) Finalize() {
	var defaultContestSlug = ""
	if ut.ContestSlug == nil {
		ut.ContestSlug = &defaultContestSlug
	}
	var defaultRanked = true
	if ut.Ranked == nil {
		ut.Ranked = &defaultRanked
	}
	var defaultTaskSlug = ""
	if ut.TaskSlug == nil {
		ut.TaskSlug = &defaultTaskSlug
	}
}

// Validate validates the entryFormPayload type instance.
func (ut *entryFormPayload) Validate() (err error) {
	if ut.ContestSlug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "contestSlug"))
	}
	if ut.TaskSlug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "taskSlug"))
	}
	if ut.Ranked == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "ranked"))
	}
	if ut.Sources == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "sources"))
	}
	if ut.ContestSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *ut.ContestSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.contestSlug`, *ut.ContestSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	if ut.Sources != nil {
		if len(ut.Sources) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.sources`, ut.Sources, len(ut.Sources), 1, true))
		}
	}
	if ut.TaskSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *ut.TaskSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.taskSlug`, *ut.TaskSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	return
}

// Publicize creates EntryFormPayload from entryFormPayload
func (ut *entryFormPayload) Publicize() *EntryFormPayload {
	var pub EntryFormPayload
	if ut.ContestSlug != nil {
		pub.ContestSlug = *ut.ContestSlug
	}
	if ut.Ranked != nil {
		pub.Ranked = *ut.Ranked
	}
	if ut.Sources != nil {
		pub.Sources = ut.Sources
	}
	if ut.TaskSlug != nil {
		pub.TaskSlug = *ut.TaskSlug
	}
	return &pub
}

// Any source code or input that should be compiled, executed or evaluated
type EntryFormPayload struct {
	// Contest unique and human readable string identifier
	ContestSlug string `form:"contestSlug" json:"contestSlug" yaml:"contestSlug" xml:"contestSlug"`
	// Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words
	// 		any submitted Entry will use a CMS Token
	Ranked bool `form:"ranked" json:"ranked" yaml:"ranked" xml:"ranked"`
	// Source files representation. Within this list the source code files and input files can be
	// 						  sent alike.
	Sources []multipart.FileHeader `form:"sources" json:"sources" yaml:"sources" xml:"sources"`
	// Task unique and human readable string identifier
	TaskSlug string `form:"taskSlug" json:"taskSlug" yaml:"taskSlug" xml:"taskSlug"`
}

// Validate validates the EntryFormPayload type instance.
func (ut *EntryFormPayload) Validate() (err error) {
	if ut.ContestSlug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "contestSlug"))
	}
	if ut.TaskSlug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "taskSlug"))
	}

	if ut.Sources == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "sources"))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, ut.ContestSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.contestSlug`, ut.ContestSlug, `[_a-zA-Z0-9\-]+`))
	}
	if len(ut.Sources) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.sources`, ut.Sources, len(ut.Sources), 1, true))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, ut.TaskSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.taskSlug`, ut.TaskSlug, `[_a-zA-Z0-9\-]+`))
	}
	return
}

// Any source code or input that should be compiled, executed or evaluated
type entryPayload struct {
	// Contest unique and human readable string identifier
	ContestSlug *string `form:"contestSlug,omitempty" json:"contestSlug,omitempty" yaml:"contestSlug,omitempty" xml:"contestSlug,omitempty"`
	// Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words
	// 		any submitted Entry will use a CMS Token
	Ranked *bool `form:"ranked,omitempty" json:"ranked,omitempty" yaml:"ranked,omitempty" xml:"ranked,omitempty"`
	// Source files representation. Within this list the source code files and input files can be sent alike.
	Sources []*entrySource `form:"sources,omitempty" json:"sources,omitempty" yaml:"sources,omitempty" xml:"sources,omitempty"`
	// Task unique and human readable string identifier
	TaskSlug *string `form:"taskSlug,omitempty" json:"taskSlug,omitempty" yaml:"taskSlug,omitempty" xml:"taskSlug,omitempty"`
}

// Finalize sets the default values for entryPayload type instance.
func (ut *entryPayload) Finalize() {
	var defaultContestSlug = ""
	if ut.ContestSlug == nil {
		ut.ContestSlug = &defaultContestSlug
	}
	var defaultRanked = true
	if ut.Ranked == nil {
		ut.Ranked = &defaultRanked
	}
	for _, e := range ut.Sources {
		var defaultEncoding = "utf8"
		if e.Encoding == nil {
			e.Encoding = &defaultEncoding
		}
	}
	var defaultTaskSlug = ""
	if ut.TaskSlug == nil {
		ut.TaskSlug = &defaultTaskSlug
	}
}

// Validate validates the entryPayload type instance.
func (ut *entryPayload) Validate() (err error) {
	if ut.ContestSlug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "contestSlug"))
	}
	if ut.TaskSlug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "taskSlug"))
	}
	if ut.Ranked == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "ranked"))
	}
	if ut.Sources == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "sources"))
	}
	if ut.ContestSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *ut.ContestSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.contestSlug`, *ut.ContestSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	if ut.Sources != nil {
		if len(ut.Sources) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.sources`, ut.Sources, len(ut.Sources), 1, true))
		}
	}
	if ut.TaskSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *ut.TaskSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.taskSlug`, *ut.TaskSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	return
}

// Publicize creates EntryPayload from entryPayload
func (ut *entryPayload) Publicize() *EntryPayload {
	var pub EntryPayload
	if ut.ContestSlug != nil {
		pub.ContestSlug = *ut.ContestSlug
	}
	if ut.Ranked != nil {
		pub.Ranked = *ut.Ranked
	}
	if ut.Sources != nil {
		pub.Sources = make([]*EntrySource, len(ut.Sources))
		for i2, elem2 := range ut.Sources {
			pub.Sources[i2] = elem2.Publicize()
		}
	}
	if ut.TaskSlug != nil {
		pub.TaskSlug = *ut.TaskSlug
	}
	return &pub
}

// Any source code or input that should be compiled, executed or evaluated
type EntryPayload struct {
	// Contest unique and human readable string identifier
	ContestSlug string `form:"contestSlug" json:"contestSlug" yaml:"contestSlug" xml:"contestSlug"`
	// Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words
	// 		any submitted Entry will use a CMS Token
	Ranked bool `form:"ranked" json:"ranked" yaml:"ranked" xml:"ranked"`
	// Source files representation. Within this list the source code files and input files can be sent alike.
	Sources []*EntrySource `form:"sources" json:"sources" yaml:"sources" xml:"sources"`
	// Task unique and human readable string identifier
	TaskSlug string `form:"taskSlug" json:"taskSlug" yaml:"taskSlug" xml:"taskSlug"`
}

// Validate validates the EntryPayload type instance.
func (ut *EntryPayload) Validate() (err error) {
	if ut.ContestSlug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "contestSlug"))
	}
	if ut.TaskSlug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "taskSlug"))
	}

	if ut.Sources == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "sources"))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, ut.ContestSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.contestSlug`, ut.ContestSlug, `[_a-zA-Z0-9\-]+`))
	}
	if len(ut.Sources) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.sources`, ut.Sources, len(ut.Sources), 1, true))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, ut.TaskSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.taskSlug`, ut.TaskSlug, `[_a-zA-Z0-9\-]+`))
	}
	return
}

// Entry's embed type which represents a source file
type entrySource struct {
	// Source content
	Content *string `form:"content,omitempty" json:"content,omitempty" yaml:"content,omitempty" xml:"content,omitempty"`
	// Source content's encoding
	Encoding *string `form:"encoding,omitempty" json:"encoding,omitempty" yaml:"encoding,omitempty" xml:"encoding,omitempty"`
	// Identifies the programming language used in the entry's content. The special keyword "none" should be used
	// 		instead when submitting plain text, which are used for user test inputs and  diff based grading
	Language *string `form:"language,omitempty" json:"language,omitempty" yaml:"language,omitempty" xml:"language,omitempty"`
	// Source file name including its extension. This field's value should comply with the name format constraint
	// 		declared by the task resource. Taking the "batch.%l" format as example, the valid source code file names could
	// 		be "batch.py", "batch.cpp" or "batch.js"
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Finalize sets the default values for entrySource type instance.
func (ut *entrySource) Finalize() {
	var defaultEncoding = "utf8"
	if ut.Encoding == nil {
		ut.Encoding = &defaultEncoding
	}
}

// Publicize creates EntrySource from entrySource
func (ut *entrySource) Publicize() *EntrySource {
	var pub EntrySource
	if ut.Content != nil {
		pub.Content = ut.Content
	}
	if ut.Encoding != nil {
		pub.Encoding = *ut.Encoding
	}
	if ut.Language != nil {
		pub.Language = ut.Language
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// Entry's embed type which represents a source file
type EntrySource struct {
	// Source content
	Content *string `form:"content,omitempty" json:"content,omitempty" yaml:"content,omitempty" xml:"content,omitempty"`
	// Source content's encoding
	Encoding string `form:"encoding" json:"encoding" yaml:"encoding" xml:"encoding"`
	// Identifies the programming language used in the entry's content. The special keyword "none" should be used
	// 		instead when submitting plain text, which are used for user test inputs and  diff based grading
	Language *string `form:"language,omitempty" json:"language,omitempty" yaml:"language,omitempty" xml:"language,omitempty"`
	// Source file name including its extension. This field's value should comply with the name format constraint
	// 		declared by the task resource. Taking the "batch.%l" format as example, the valid source code file names could
	// 		be "batch.py", "batch.cpp" or "batch.js"
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Embedded reprensentation of an entry evaluation result
type evaluationResult struct {
	// Execution result status
	Status *string `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
	// Evaluation retries
	Tries *int `form:"tries,omitempty" json:"tries,omitempty" yaml:"tries,omitempty" xml:"tries,omitempty"`
}

// Finalize sets the default values for evaluationResult type instance.
func (ut *evaluationResult) Finalize() {
	var defaultStatus = "unprocessed"
	if ut.Status == nil {
		ut.Status = &defaultStatus
	}
}

// Validate validates the evaluationResult type instance.
func (ut *evaluationResult) Validate() (err error) {
	if ut.Status != nil {
		if !(*ut.Status == "ok" || *ut.Status == "unprocessed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.status`, *ut.Status, []interface{}{"ok", "unprocessed"}))
		}
	}
	if ut.Tries != nil {
		if *ut.Tries < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.tries`, *ut.Tries, 0, true))
		}
	}
	return
}

// Publicize creates EvaluationResult from evaluationResult
func (ut *evaluationResult) Publicize() *EvaluationResult {
	var pub EvaluationResult
	if ut.Status != nil {
		pub.Status = *ut.Status
	}
	if ut.Tries != nil {
		pub.Tries = ut.Tries
	}
	return &pub
}

// Embedded reprensentation of an entry evaluation result
type EvaluationResult struct {
	// Execution result status
	Status string `form:"status" json:"status" yaml:"status" xml:"status"`
	// Evaluation retries
	Tries *int `form:"tries,omitempty" json:"tries,omitempty" yaml:"tries,omitempty" xml:"tries,omitempty"`
}

// Validate validates the EvaluationResult type instance.
func (ut *EvaluationResult) Validate() (err error) {
	if !(ut.Status == "ok" || ut.Status == "unprocessed") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.status`, ut.Status, []interface{}{"ok", "unprocessed"}))
	}
	if ut.Tries != nil {
		if *ut.Tries < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.tries`, *ut.Tries, 0, true))
		}
	}
	return
}

// Embedded representation of an entry execution result
type executionResult struct {
	// Memory consumed
	Memory *int `form:"memory,omitempty" json:"memory,omitempty" yaml:"memory,omitempty" xml:"memory,omitempty"`
	// Execution result status
	Status *string `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
	// The spent execution CPU time
	Time *float64 `form:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty" xml:"time,omitempty"`
	// The spent execution human perceived time
	WallClockTime *float64 `form:"wallClockTime,omitempty" json:"wallClockTime,omitempty" yaml:"wallClockTime,omitempty" xml:"wallClockTime,omitempty"`
}

// Publicize creates ExecutionResult from executionResult
func (ut *executionResult) Publicize() *ExecutionResult {
	var pub ExecutionResult
	if ut.Memory != nil {
		pub.Memory = ut.Memory
	}
	if ut.Status != nil {
		pub.Status = ut.Status
	}
	if ut.Time != nil {
		pub.Time = ut.Time
	}
	if ut.WallClockTime != nil {
		pub.WallClockTime = ut.WallClockTime
	}
	return &pub
}

// Embedded representation of an entry execution result
type ExecutionResult struct {
	// Memory consumed
	Memory *int `form:"memory,omitempty" json:"memory,omitempty" yaml:"memory,omitempty" xml:"memory,omitempty"`
	// Execution result status
	Status *string `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
	// The spent execution CPU time
	Time *float64 `form:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty" xml:"time,omitempty"`
	// The spent execution human perceived time
	WallClockTime *float64 `form:"wallClockTime,omitempty" json:"wallClockTime,omitempty" yaml:"wallClockTime,omitempty" xml:"wallClockTime,omitempty"`
}
