package model

// Type ...
type Type int

const (

	// FS Finish-to-Start – the first task needs to be completed before
	//  the second task can start, as per the example above
	FS Type = iota

	// FF Finish-to-Finish – the second task can’t be completed until
	//  the first task has been done. For example, wires can’t be
	//  fitted into the wall until they’ve been inspected.
	FF

	//SF Start-to-Finish – the first task has to start before the second
	//  task can be completed. For example, a new software installation
	//  has to start before the old installation can be stopped.
	SF

	//SS Start-to-Start – the following task can’t commence until the
	// first task has started. For example, a concrete floor can’t
	//  start to be levelled until the concrete has started pouring
	//  into the designated space. ...
	SS
)
