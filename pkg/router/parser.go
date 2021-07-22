package router

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"regexp"
)

// Function is used to register route
type Function struct {
	// route function name.
	Name string
	// route method, also is HTTP method,
	// 'MethodType' defined in method.go.
	Method MethodType
	// url of route.
	Url string
	// reciever of the Controller method.
	Recv reflect.Value
	// dto will be automatically instantiated according to the body
	// in json format.
	Dto reflect.Value
}

type Parser struct {
	Funcs   []*Function
	visitor *visitor
}

// Parse registered controllers to a parser,
// parser contains function slice, each function
// has the information needed by the router.
func NewParser(controllers []interface{}) *Parser {
	dir, _ := os.Getwd()
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir+"/controllers", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	parser := &Parser{
		make([]*Function, 0, len(controllers)),
		&visitor{make(map[string]*visitFunc)},
	}

	if pkg, ok := pkgs["controllers"]; ok {
		parsePackage(pkg, parser.visitor)
	}

	for _, controller := range controllers {
		ct := reflect.TypeOf(controller)
		for i := 0; i < ct.NumMethod(); i++ {
			method := ct.Method(i)
			parser.methodToFunc(method, controller)
		}
	}

	return parser
}

// According ast result and reflect method to append a function to parser
func (p *Parser) methodToFunc(method reflect.Method, controller interface{}) {
	if v, ok := p.visitor.funcs[method.Type.In(0).Elem().Name()+method.Name]; ok {
		httpMethod, url := annotateComment(v.comment)
		if httpMethod == Unknown || url == "" {
			// not a route function that meets the matching rules
			return
		}

		var dto reflect.Value
		if method.Type.NumIn() > 2 {
			dto = reflect.New(method.Type.In(2)).Elem()
		}

		p.Funcs = append(p.Funcs, &Function{
			Name:   method.Name,
			Method: httpMethod,
			Url:    url,
			Recv:   reflect.ValueOf(controller),
			Dto:    dto,
		})

	}
}

// Parse a package
func parsePackage(ap *ast.Package, v ast.Visitor) {
	for _, file := range ap.Files {
		ast.Walk(v, file)
	}
}

// Annotate function's comment to get method and url
func annotateComment(comment string) (MethodType, string) {
	commentReg := regexp.MustCompile(`\[(.*?)\(\"(.*?)\"\)\]`)
	if commentReg == nil {
		return Unknown, ""
	}
	res := commentReg.FindStringSubmatch(comment)
	if len(res) == 3 {
		return getMethodType(res[1]), res[2]
	}
	return Unknown, ""
}

// Visitor visit ast node, save function info to the funcs
type visitor struct {
	// map key is controller's name + function's name,
	// value is comment from ast.
	funcs map[string]*visitFunc
}

type visitFunc struct {
	comment  string
	inNames  []string
	outNames []string
}

// Implement ast.Visitor to parse func node
func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node != nil {
		if funcDecl, ok := node.(*ast.FuncDecl); ok && isValidFunc(funcDecl) {
			var typeName string
			ft := funcDecl.Recv.List[0].Type

			if starExpr, ok := ft.(*ast.StarExpr); ok {
				// method reciever is point.
				//ex:(x *X)
				typeName = starExpr.X.(*ast.Ident).Obj.Name
			} else if ident, ok := ft.(*ast.Ident); ok {
				// method reciever is type struct.
				//ex:(x X)
				typeName = ident.Name
			} else {
				return
			}

			inNames := make([]string, 0)
			outNames := make([]string, 0)
			if funcDecl.Type.Params != nil {
				for _, i := range funcDecl.Type.Params.List {
					inNames = append(inNames, i.Names[0].Name)
				}
			}
			if funcDecl.Type.Results != nil {
				for _, i := range funcDecl.Type.Results.List {
					outNames = append(outNames, i.Names[0].Name)
				}
			}

			v.funcs[typeName+funcDecl.Name.Name] = &visitFunc{
				comment:  funcDecl.Doc.Text(),
				inNames:  inNames,
				outNames: outNames,
			}
		}
	}
	return v
}

func isValidFunc(funcD *ast.FuncDecl) bool {
	if funcD.Recv == nil || funcD.Recv.List == nil || len(funcD.Recv.List) == 0 {
		return false
	}
	return true
}
