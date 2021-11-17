package ast

import (
	parsec "github.com/prataprc/goparsec"
)

func makey(ast *parsec.AST) parsec.Parser {
	slash := parsec.Atom(`/`, "SLASH")
	meterid := parsec.Token(`[a-zA-Z0-9.\-_,:\x02\\ ]+`, "METERID")
	parenopen := parsec.Atom(`(`, "PARENOPEN")
	parenclose := parsec.Atom(`)`, "PARENCLOSE")
	starsep := parsec.Atom(`*`, "STARSEP")
	unit := parsec.Token(`[a-zA-Z0-9.\-_]+`, "UNIT")
	value := parsec.Token(`[+\-]?[a-zA-Z0-9.\-_]+`, "VALUE")
	intv := parsec.Token(`[0-9]+`, "INT")
	timestamp := parsec.Token(`[0-9]{12}[SW]?`, "TIMESTAMP")
	exclamation := parsec.Atom(`!`, "EXCLAMATION")
	checksum1 := parsec.Token(`.`, "CHECKSUM1")
	checksum4 := parsec.Token(`.{4}`, "CHECKSUM4")
	id := parsec.Token(`([0-9]+-[0-9]+:[0-9]+\.)?[0-9]+.[0-9]+(\*[0-9]+)?`, "ID")

	identification := ast.And("identification", nil, slash, meterid)

	cosemempty := ast.And("cosemempty", nil, parenopen, parenclose)
	cosemunitless := ast.And("cosemunitless", nil, parenopen, value, parenclose)
	cosemwithunit := ast.And("cosemwithunit", nil, parenopen, value, starsep, unit, parenclose)
	cosemvalue := ast.OrdChoice("cosemvalue", nil, cosemunitless, cosemempty, cosemwithunit)

	logentry := ast.And("logentry", nil, parenopen, timestamp, parenclose, cosemvalue)
	logentries := ast.Kleene("logentries", nil, logentry)

	cosem := ast.And("cosem", nil, id, cosemvalue)
	mbus := ast.And("mbus", nil, id, parenopen, timestamp, parenclose, cosemvalue)
	dsmr3gas := ast.And("dsmr3gas", nil, id, parenopen, timestamp, parenclose, parenopen, intv, parenclose, parenopen, intv, parenclose, parenopen, intv, parenclose, parenopen, id, parenclose, parenopen, unit, parenclose, parenopen, value, parenclose)
	profilegeneric := ast.And("profilegeneric", nil, id, parenopen, intv, parenclose, parenopen, id, parenclose, logentries)

	object := ast.OrdChoice("object", nil, dsmr3gas, profilegeneric, mbus, cosem)

	objects := ast.Kleene("objects", nil, object)

	checksum := ast.Maybe("optchecksum", nil, ast.OrdChoice("checksum", nil, checksum4, checksum1))

	footer := ast.And("footer", nil, exclamation, checksum)
	return ast.And("telegram", nil, identification, objects, footer)
}
