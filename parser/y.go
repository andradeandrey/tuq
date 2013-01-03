
//line unql.y:2
package parser
import "fmt"
import "log"

func logDebugGrammar(format string, v ...interface{}) {
    if debugGrammar && len(v) > 0 {
        log.Printf("DEBUG GRAMMAR " + format, v)
    } else if debugGrammar {
        log.Printf("DEBUG GRAMMAR " + format)
    }
}

//line unql.y:15
type yySymType struct {
	yys int 
s string 
n int
f float64}

const INT = 57346
const REAL = 57347
const STRING = 57348
const IDENTIFIER = 57349
const PROPERTY = 57350
const NEWLINE = 57351
const LPAREN = 57352
const RPAREN = 57353
const COMMA = 57354
const LBRACE = 57355
const RBRACE = 57356
const SELECT = 57357
const DISTINCT = 57358
const ALL = 57359
const AS = 57360
const FROM = 57361
const WHERE = 57362
const GROUP = 57363
const BY = 57364
const HAVING = 57365
const UNION = 57366
const INTERSECT = 57367
const EXCEPT = 57368
const ORDER = 57369
const LIMIT = 57370
const OFFSET = 57371
const ASC = 57372
const DESC = 57373
const TRUE = 57374
const FALSE = 57375
const LBRACKET = 57376
const RBRACKET = 57377
const QUESTION = 57378
const COLON = 57379
const MAX = 57380
const MIN = 57381
const AVG = 57382
const COUNT = 57383
const SUM = 57384
const DOT = 57385
const PLUS = 57386
const MINUS = 57387
const MULT = 57388
const DIV = 57389
const MOD = 57390
const AND = 57391
const OR = 57392
const NOT = 57393
const EQ = 57394
const LT = 57395
const LTE = 57396
const GT = 57397
const GTE = 57398
const NE = 57399
const PRAGMA = 57400
const ASSIGN = 57401
const EXPLAIN = 57402
const NULL = 57403
const OVER = 57404

var yyToknames = []string{
	"INT",
	"REAL",
	"STRING",
	"IDENTIFIER",
	"PROPERTY",
	"NEWLINE",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"LBRACE",
	"RBRACE",
	"SELECT",
	"DISTINCT",
	"ALL",
	"AS",
	"FROM",
	"WHERE",
	"GROUP",
	"BY",
	"HAVING",
	"UNION",
	"INTERSECT",
	"EXCEPT",
	"ORDER",
	"LIMIT",
	"OFFSET",
	"ASC",
	"DESC",
	"TRUE",
	"FALSE",
	"LBRACKET",
	"RBRACKET",
	"QUESTION",
	"COLON",
	"MAX",
	"MIN",
	"AVG",
	"COUNT",
	"SUM",
	"DOT",
	"PLUS",
	"MINUS",
	"MULT",
	"DIV",
	"MOD",
	"AND",
	"OR",
	"NOT",
	"EQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"NE",
	"PRAGMA",
	"ASSIGN",
	"EXPLAIN",
	"NULL",
	"OVER",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line unql.y:390


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 97
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 200

var yyAct = []int{

	71, 124, 70, 116, 67, 19, 12, 84, 127, 48,
	75, 45, 125, 5, 128, 6, 106, 20, 22, 23,
	30, 129, 107, 29, 64, 114, 26, 7, 136, 137,
	73, 78, 37, 139, 13, 50, 51, 52, 53, 62,
	39, 40, 41, 140, 79, 24, 25, 27, 2, 88,
	89, 32, 31, 33, 34, 35, 127, 121, 21, 50,
	51, 52, 53, 62, 15, 103, 80, 56, 57, 58,
	59, 60, 61, 6, 18, 109, 83, 43, 74, 115,
	118, 112, 150, 81, 122, 90, 91, 92, 93, 94,
	95, 96, 97, 98, 99, 100, 101, 102, 47, 46,
	14, 11, 104, 135, 123, 108, 105, 131, 133, 111,
	130, 132, 110, 72, 151, 134, 63, 20, 22, 23,
	30, 30, 142, 29, 87, 69, 26, 68, 143, 145,
	28, 141, 17, 144, 65, 66, 118, 16, 44, 146,
	147, 10, 126, 148, 149, 24, 25, 27, 86, 85,
	138, 32, 31, 33, 34, 35, 120, 119, 21, 49,
	82, 42, 9, 38, 15, 8, 117, 50, 51, 52,
	53, 62, 54, 55, 18, 56, 57, 58, 59, 60,
	61, 50, 51, 52, 53, 62, 54, 113, 77, 56,
	57, 58, 59, 60, 61, 76, 36, 4, 3, 1,
}
var yyPact = []int{

	-45, -1000, -1000, -1000, 86, 113, -1000, 5, 16, 58,
	113, 82, -50, 123, -1000, 113, -1000, -1000, -1000, -10,
	-1000, 130, -1000, -1000, -1000, -1000, 119, 113, 103, 13,
	-33, -1000, -1000, -1000, -1000, -1000, 3, 22, 86, 66,
	-1000, -1000, 56, 117, -1000, -1000, -1000, -1000, 113, 113,
	113, 113, 113, 113, 113, 113, 113, 113, 113, 113,
	113, 113, 113, -1000, 113, -1000, -1000, 88, 94, -21,
	-13, 93, 113, 101, 98, 114, -1000, -4, 113, 113,
	-1000, -1000, 36, 113, -1000, 92, -6, -1000, -1000, -23,
	-1000, -1000, -1000, -1000, 15, 137, -9, -9, -9, -9,
	-9, -9, -1000, -14, -1000, 119, 113, -1000, 113, 97,
	-1000, -1000, -1000, -1000, 113, -1000, -1000, 91, -2, -1000,
	10, 21, -1000, 117, -1000, 115, -54, 114, 113, -1000,
	-1000, -1000, -1000, -1000, -1000, 113, -1000, -1000, -1000, 113,
	113, -1000, -54, -1000, 64, -1000, -1000, -1000, -1000, -1000,
	107, -1000,
}
var yyPgo = []int{

	0, 199, 48, 198, 0, 197, 27, 196, 195, 3,
	188, 187, 166, 165, 163, 162, 161, 160, 157, 156,
	2, 150, 7, 149, 148, 1, 142, 5, 141, 138,
	34, 100, 137, 132, 4, 130, 127,
}
var yyR1 = []int{

	0, 1, 1, 3, 2, 5, 5, 7, 7, 8,
	8, 8, 10, 11, 9, 9, 12, 12, 12, 6,
	6, 14, 14, 14, 14, 13, 19, 18, 18, 18,
	21, 17, 17, 16, 16, 22, 22, 23, 23, 23,
	23, 25, 25, 26, 24, 15, 28, 28, 28, 29,
	29, 4, 4, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 31, 31, 32,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 20, 20, 34, 34, 36,
	27, 27, 35, 35, 35, 35, 35,
}
var yyR2 = []int{

	0, 1, 1, 4, 4, 0, 1, 0, 3, 0,
	1, 2, 2, 2, 1, 3, 1, 2, 2, 1,
	3, 1, 2, 1, 1, 4, 3, 0, 1, 2,
	2, 0, 2, 0, 2, 1, 3, 1, 2, 3,
	4, 1, 2, 4, 1, 2, 1, 2, 2, 0,
	1, 1, 5, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 2, 1, 1,
	1, 1, 4, 1, 2, 1, 2, 1, 1, 1,
	3, 3, 4, 3, 3, 1, 3, 1, 3, 3,
	1, 3, 1, 1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -5, 58, 60, -6, -13, -15,
	-28, 15, -4, -30, -31, 51, -32, -33, 61, -27,
	4, 45, 5, 6, 32, 33, 13, 34, -35, 10,
	7, 39, 38, 40, 41, 42, -7, 27, -14, 24,
	25, 26, -16, 19, -29, -4, 17, 16, 59, 36,
	44, 45, 46, 47, 49, 50, 52, 53, 54, 55,
	56, 57, 48, -31, 34, 4, 5, -34, -36, 6,
	-20, -4, 10, -4, -2, 43, -8, -10, 28, 22,
	-6, 17, -17, 20, -22, -23, -24, 7, -4, -4,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -4, 14, 12, 37, 35, 12, -20,
	11, 11, -27, -11, 29, -4, -9, -12, -4, -18,
	-19, 21, -4, 12, -25, 18, -26, 62, 37, 35,
	-34, -4, -20, 11, -4, 12, 30, 31, -21, 23,
	22, -22, 7, -25, -27, -4, -9, -4, -20, -25,
	18, 7,
}
var yyDef = []int{

	5, -2, 1, 2, 0, 0, 6, 7, 19, 33,
	49, 46, 0, 51, 66, 0, 68, 69, 70, 71,
	73, 0, 75, 77, 78, 79, 0, 0, 0, 5,
	90, 92, 93, 94, 95, 96, 9, 0, 0, 21,
	23, 24, 31, 0, 45, 50, 47, 48, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 67, 0, 74, 76, 0, 87, 0,
	0, 85, 0, 0, 0, 0, 4, 10, 0, 0,
	20, 22, 27, 0, 34, 35, 37, 44, 3, 0,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 65, 0, 80, 0, 0, 81, 0, 0,
	83, 84, 91, 11, 0, 12, 8, 14, 16, 25,
	28, 0, 32, 0, 38, 0, 41, 0, 0, 72,
	88, 89, 86, 82, 13, 0, 17, 18, 29, 0,
	0, 36, 39, 42, 0, 52, 15, 30, 26, 40,
	0, 43,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c > 0 && c <= len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return fmt.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return fmt.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		fmt.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		fmt.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				fmt.Printf("%s", yyStatname(yystate))
				fmt.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					fmt.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				fmt.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		fmt.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line unql.y:35
		{ logDebugGrammar("INPUT") }
	case 3:
		//line unql.y:39
		{ logDebugGrammar("PRAGMA: %v", yyS[yypt-3])
	                                                       right := parsingStack.Pop()
	                                                       left := parsingStack.Pop()
	                                                       ProcessPragma(left.(Expression), right.(Expression))
	                                                     }
	case 4:
		//line unql.y:46
		{ logDebugGrammar("SELECT_STMT")
	                                                                    parsingQuery.parsedSuccessfully = true }
	case 6:
		//line unql.y:51
		{  parsingQuery.isExplainOnly = true  }
	case 12:
		//line unql.y:63
		{ thisExpression := parsingStack.Pop()
	                                   parsingQuery.Limit = thisExpression.(Expression)
	                                 }
	case 13:
		//line unql.y:68
		{ thisExpression := parsingStack.Pop()
	                                   parsingQuery.Offset = thisExpression.(Expression)
	                                 }
	case 16:
		//line unql.y:77
		{ thisExpression := NewSortItem(parsingStack.Pop().(Expression), true)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 17:
		//line unql.y:80
		{ thisExpression := NewSortItem(parsingStack.Pop().(Expression), true)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 18:
		//line unql.y:83
		{ thisExpression := NewSortItem(parsingStack.Pop().(Expression), false)
	                            parsingQuery.Orderby = append(parsingQuery.Orderby, *thisExpression)
	                           }
	case 19:
		//line unql.y:88
		{ logDebugGrammar("SELECT_COMPOUND") }
	case 25:
		//line unql.y:98
		{ logDebugGrammar("SELECT_CORE") }
	case 26:
		//line unql.y:101
		{ logDebugGrammar("SELECT GROUP")
	                                         parsingQuery.isAggregateQuery = true 
	                                         parsingQuery.Groupby = parsingStack.Pop().(ExpressionList) }
	case 27:
		//line unql.y:106
		{ logDebugGrammar("SELECT GROUP HAVING - EMPTY") }
	case 28:
		//line unql.y:107
		{ logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP") }
	case 29:
		//line unql.y:108
		{ logDebugGrammar("SELECT GROUP HAVING - SELECT GROUP SELECT HAVING") }
	case 30:
		//line unql.y:113
		{ parsingQuery.Having = parsingStack.Pop().(Expression) }
	case 31:
		//line unql.y:116
		{ logDebugGrammar("SELECT WHERE - EMPTY") }
	case 32:
		//line unql.y:117
		{ logDebugGrammar("SELECT WHERE - EXPR")
	                               where_part := parsingStack.Pop()
	                               parsingQuery.Where = where_part.(Expression) }
	case 34:
		//line unql.y:123
		{ logDebugGrammar("SELECT_FROM") }
	case 37:
		//line unql.y:130
		{ ds := NewDataSource(yyS[yypt-0].s)
	                                   parsingQuery.AddDataSource(ds) 
	                                 }
	case 38:
		//line unql.y:133
		{ ds := NewDataSource(yyS[yypt-1].s)
	                                          nextOver := parsingStack.Pop()
	                                          for nextOver != nil {
	                                            ds.AddOver(nextOver.(*Over))
	                                            nextOver = parsingStack.Pop()
	                                          }
	                                          parsingQuery.AddDataSource(ds)
	                                        }
	case 39:
		//line unql.y:141
		{ ds := NewDataSourceWithAs(yyS[yypt-2].s, yyS[yypt-0].s) 
	                                          parsingQuery.AddDataSource(ds) 
	                                        }
	case 40:
		//line unql.y:144
		{ ds := NewDataSourceWithAs(yyS[yypt-3].s, yyS[yypt-1].s)
	                                          nextOver := parsingStack.Pop()
	                                          for nextOver != nil {
	                                            ds.AddOver(nextOver.(*Over))
	                                            nextOver = parsingStack.Pop()
	                                          }
	                                          parsingQuery.AddDataSource(ds)
	                                        }
	case 43:
		//line unql.y:158
		{   prop := parsingStack.Pop().(*Property)
	                                                    over := NewOver(prop, yyS[yypt-0].s)
	                                                    parsingStack.Push(over)
	                                                  }
	case 45:
		//line unql.y:166
		{ logDebugGrammar("SELECT_SELECT") }
	case 46:
		//line unql.y:169
		{ logDebugGrammar("SELECT_SELECT_HEAD") }
	case 49:
		//line unql.y:174
		{ logDebugGrammar("SELECT SELECT TAIL - EMPTY") }
	case 50:
		//line unql.y:175
		{ logDebugGrammar("SELECT SELECT TAIL - EXPR")
	                            thisExpression := parsingStack.Pop()
	                            parsingQuery.Sel = thisExpression.(Expression)
	                          }
	case 51:
		//line unql.y:181
		{ logDebugGrammar("EXPRESSION") }
	case 52:
		//line unql.y:182
		{ logDebugGrammar("EXPRESSION - TERNARY")
	                                                    elsee := parsingStack.Pop().(Expression)
	                                                    thenn := parsingStack.Pop().(Expression)
	                                                    iff := parsingStack.Pop().(Expression)
	                                                    thisExpr := NewTernaryExpression(iff, thenn, elsee)
	                                                    parsingStack.Push(thisExpr)
	                                                  }
	case 53:
		//line unql.y:191
		{  logDebugGrammar("EXPR - PLUS")
	                        right := parsingStack.Pop()
	                        left := parsingStack.Pop()
	                        thisExpression := NewPlusExpression(left.(Expression), right.(Expression)) 
	                        parsingStack.Push(thisExpression)
	                     }
	case 54:
		//line unql.y:197
		{  logDebugGrammar("EXPR - MINUS")
	                               right := parsingStack.Pop()
	                               left := parsingStack.Pop()
	                               thisExpression := NewMinusExpression(left.(Expression), right.(Expression)) 
	                               parsingStack.Push(thisExpression)
	                            }
	case 55:
		//line unql.y:203
		{  logDebugGrammar("EXPR - MULT")
	                              right := parsingStack.Pop()
	                              left := parsingStack.Pop()
	                              thisExpression := NewMultiplyExpression(left.(Expression), right.(Expression)) 
	                              parsingStack.Push(thisExpression)
	                           }
	case 56:
		//line unql.y:209
		{  logDebugGrammar("EXPR - DIV")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewDivideExpression(left.(Expression), right.(Expression)) 
	                             parsingStack.Push(thisExpression)
	                          }
	case 57:
		//line unql.y:215
		{  logDebugGrammar("EXPR - AND")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewAndExpression(left.(Expression), right.(Expression)) 
	                             parsingStack.Push(thisExpression)
	                         }
	case 58:
		//line unql.y:221
		{  logDebugGrammar("EXPR - OR")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewOrExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 59:
		//line unql.y:227
		{  logDebugGrammar("EXPR - EQ")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewEqualsExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 60:
		//line unql.y:233
		{  logDebugGrammar("EXPR - LT")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewLessThanExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 61:
		//line unql.y:239
		{  logDebugGrammar("EXPR - LTE")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewLessThanOrEqualExpression(left.(Expression), right.(Expression)) 
	                             parsingStack.Push(thisExpression)
	                         }
	case 62:
		//line unql.y:245
		{  logDebugGrammar("EXPR - GT")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewGreaterThanExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 63:
		//line unql.y:251
		{  logDebugGrammar("EXPR - GTE")
	                             right := parsingStack.Pop()
	                             left := parsingStack.Pop()
	                             thisExpression := NewGreaterThanOrEqualExpression(left.(Expression), right.(Expression)) 
	                             parsingStack.Push(thisExpression)
	                         }
	case 64:
		//line unql.y:257
		{  logDebugGrammar("EXPR - NE")
	                            right := parsingStack.Pop()
	                            left := parsingStack.Pop()
	                            thisExpression := NewNotEqualsExpression(left.(Expression), right.(Expression)) 
	                            parsingStack.Push(thisExpression)
	                         }
	case 67:
		//line unql.y:268
		{ logDebugGrammar("EXPR - NOT")
	                               curr := parsingStack.Pop().(Expression)
	                               thisExpression := NewNotExpression(curr)
	                               parsingStack.Push(thisExpression)
	                             }
	case 69:
		//line unql.y:276
		{ logDebugGrammar("SUFFIX_EXPR") }
	case 70:
		//line unql.y:279
		{ logDebugGrammar("NULL")
	             thisExpression := NewNull()
	             parsingStack.Push(thisExpression)
	           }
	case 71:
		//line unql.y:283
		{  }
	case 72:
		//line unql.y:284
		{     logDebugGrammar("ATOM - prop[]")
	                                                  rightExpr := parsingStack.Pop().(Expression)
	                                                  leftProp := parsingStack.Pop().(*Property)
	                                                  thisExpression := NewBracketMemberExpression(leftProp, rightExpr)
	                                                  parsingStack.Push(thisExpression)
	                                            }
	case 73:
		//line unql.y:290
		{ thisExpression := NewIntegerLiteral(yyS[yypt-0].n) 
	                 parsingStack.Push(thisExpression) }
	case 74:
		//line unql.y:292
		{ thisExpression := NewIntegerLiteral(-yyS[yypt-1].n)
	                 parsingStack.Push(thisExpression) }
	case 75:
		//line unql.y:294
		{ thisExpression := NewFloatLiteral(yyS[yypt-0].f) 
	                 parsingStack.Push(thisExpression) }
	case 76:
		//line unql.y:296
		{ thisExpression := NewFloatLiteral(-yyS[yypt-1].f)
	                 parsingStack.Push(thisExpression) }
	case 77:
		//line unql.y:298
		{ thisExpression := NewStringLiteral(yyS[yypt-0].s) 
	                 parsingStack.Push(thisExpression) }
	case 78:
		//line unql.y:300
		{ thisExpression := NewBoolLiteral(true) 
	                 parsingStack.Push(thisExpression) }
	case 79:
		//line unql.y:302
		{ thisExpression := NewBoolLiteral(false) 
	                 parsingStack.Push(thisExpression)}
	case 80:
		//line unql.y:304
		{ logDebugGrammar("ATOM - {}")
	                                            }
	case 81:
		//line unql.y:306
		{ logDebugGrammar("ATOM - []")
	                                            exp_list := parsingStack.Pop().(ExpressionList)
	                                            thisExpression := NewArrayLiteral(exp_list)
	                                            parsingStack.Push(thisExpression)
	                                          }
	case 82:
		//line unql.y:311
		{ logDebugGrammar("FUNCTION - $1.s")
	                                                      exp_list := parsingStack.Pop().(ExpressionList)
	                                                      function := parsingStack.Pop().(*Function)
	                                                      function.AddArguments(exp_list)
	                                                      parsingStack.Push(function)
	                                                    }
	case 85:
		//line unql.y:321
		{ logDebugGrammar("EXPRESSION_LIST - EXPRESSION")
	                               exp_list := make(ExpressionList, 0)
	                               exp_list = append(exp_list, parsingStack.Pop().(Expression))
	                               parsingStack.Push(exp_list)
	                             }
	case 86:
		//line unql.y:326
		{ logDebugGrammar("EXPRESSION_LIST - EXPRESSION COMMA EXPRESSION_LIST")
	                                               rest := parsingStack.Pop().(ExpressionList)
	                                               last := parsingStack.Pop()
	                                               new_list := make(ExpressionList, 0)
	                                               new_list = append(new_list, last.(Expression))
	                                               for _, v := range rest {
	                                                new_list = append(new_list, v)
	                                               }
	                                               parsingStack.Push(new_list)
	                                             }
	case 88:
		//line unql.y:339
		{ last := parsingStack.Pop().(*ObjectLiteral)
	                                                                  rest := parsingStack.Pop().(*ObjectLiteral)
	                                                                  rest.AddAll(last)
	                                                                  parsingStack.Push(rest)
	                                                                }
	case 89:
		//line unql.y:346
		{ thisKey := yyS[yypt-2].s
	                                                     thisValue := parsingStack.Pop().(Expression)
	                                                     thisExpression := NewObjectLiteral(Object{thisKey: thisValue})
	                                                     parsingStack.Push(thisExpression) 
	                                                   }
	case 90:
		//line unql.y:353
		{
	                         thisExpression := NewProperty(yyS[yypt-0].s) 
	                         parsingStack.Push(thisExpression) 
	                       }
	case 91:
		//line unql.y:357
		{
	                                    thisValue := parsingStack.Pop().(*Property)
	                                    thisExpression := NewProperty(yyS[yypt-2].s + "." + thisValue.Symbol)
	                                    parsingStack.Push(thisExpression)
	                                  }
	case 92:
		//line unql.y:364
		{ 
	                     parsingQuery.isAggregateQuery = true
	                     thisExpression := NewFunction("min")
	                     parsingStack.Push(thisExpression)
	                   }
	case 93:
		//line unql.y:369
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("max")
	                  parsingStack.Push(thisExpression)
	                }
	case 94:
		//line unql.y:374
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("avg")
	                  parsingStack.Push(thisExpression)
	                }
	case 95:
		//line unql.y:379
		{ 
	                   parsingQuery.isAggregateQuery = true
	                   thisExpression := NewFunction("count")
	                   parsingStack.Push(thisExpression)
	                  }
	case 96:
		//line unql.y:384
		{ 
	                  parsingQuery.isAggregateQuery = true
	                  thisExpression := NewFunction("sum")
	                  parsingStack.Push(thisExpression)
	                }
	}
	goto yystack /* stack new state and value */
}
