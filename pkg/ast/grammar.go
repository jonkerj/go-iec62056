package ast

import (
	parsec "github.com/prataprc/goparsec"
)

func makey() parsec.Parser {
	// Terminals
	parenClose := parsec.Atom(`)`, "PARENCLOSE")
	parenOpen := parsec.Atom(`(`, "PARENOPEN")
	slash := parsec.Atom(`/`, "SLASH")
	starSep := parsec.Atom(`*`, "STARSEP")

	checksum1 := parsec.Token(`.`, "CHECKSUM1")
	checksum4 := parsec.Token(`.{4}`, "CHECKSUM4")
	exclamation := parsec.Atom(`!`, "EXCLAMATION")
	id := parsec.Token(`([0-9]+-[0-9]+:[0-9]+\.)?[0-9]+.[0-9]+(\*[0-9]+)?`, "ID")
	intV := parsec.Token(`[0-9]+`, "INT")
	meterID := parsec.Token(`[a-zA-Z0-9.\-_,:\x02\\ ]+`, "METERID")
	timestamp := parsec.Token(`[0-9]{12}[SW]?`, "TIMESTAMP")
	unit := parsec.Token(`[a-zA-Z0-9.\-_]+`, "UNIT")
	value := parsec.Token(`[+\-]?[a-zA-Z0-9.\-_]+`, "VALUE")

	// Non-terminals
	identification := parsec.And(nodifyIdentification, slash, meterID)

	checksum := parsec.Maybe(nodifyChecksum, parsec.OrdChoice(nodifyFirstItem, checksum4, checksum1))

	footer := parsec.And(nodifySecondItem, exclamation, checksum)

	cosemEmpty := parsec.And(nodifyCosemEmpty, parenOpen, parenClose)
	cosemUnitless := parsec.And(nodifyCosemValue, parenOpen, value, parenClose)
	cosemWithUnit := parsec.And(nodifyCosemValueUnit, parenOpen, value, starSep, unit, parenClose)
	cosemValue := parsec.OrdChoice(nodifyFirstItem, cosemUnitless, cosemEmpty, cosemWithUnit)

	logEntry := parsec.And(nil, parenOpen, timestamp, parenClose, cosemValue)
	logEntries := parsec.Kleene(nil, logEntry)

	cosem := parsec.And(nodifyCosem, id, cosemValue)
	mbus := parsec.And(nodifyMBus, id, parenOpen, timestamp, parenClose, cosemValue)
	dsmr3gas := parsec.And(nodifyIDOnly, id, parenOpen, timestamp, parenClose, parenOpen, intV, parenClose, parenOpen, intV, parenClose, parenOpen, intV, parenClose, parenOpen, id, parenClose, parenOpen, unit, parenClose, parenOpen, value, parenClose)
	profilegeneric := parsec.And(nodifyIDOnly, id, parenOpen, intV, parenClose, parenOpen, id, parenClose, logEntries)

	object := parsec.OrdChoice(nodifyFirstItem, dsmr3gas, profilegeneric, mbus, cosem)
	objects := parsec.Kleene(nil, object)

	return parsec.And(nodifyTelegram, identification, objects, footer)
}
