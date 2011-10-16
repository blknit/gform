package gform

type GeneralEventHandler func(sender Controller)

type MouseEventHandlerA func(sender Controller, arg *MouseEventArg)

type DropFilesEventHandlerA func(sender Controller, arg *DropFilesEventArg)

type PaintEventHandlerA func(sender Controller, arg *PaintEventArg)

type KeyUpEventHandlerA func(sender Controller, arg *KeyUpEventArg)

type LVEndLabelEditEventHandlerA func(sender *ListView, arg *LVEndLabelEditEventArg)

type LVDBLClickEventHandlerA func(sender *ListView, arg *LVDBLClickEventArg)