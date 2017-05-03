// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/akalenda/GolangRevelRBAC/app"
	controllers "github.com/akalenda/GolangRevelRBAC/app/controllers"
	_ "github.com/akalenda/GolangRevelRBAC/app/jobs"
	models "github.com/akalenda/GolangRevelRBAC/app/models"
	tests "github.com/akalenda/GolangRevelRBAC/tests"
	_ "github.com/mattn/go-sqlite3"
	controllers2 "github.com/revel/modules/jobs/app/controllers"
	_ "github.com/revel/modules/jobs/app/jobs"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.GorpController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Begin",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Commit",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Rollback",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					129: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Jobs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Status",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					28: []string{ 
						"entries",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Application)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "AddUser",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					38: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					42: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "SaveUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "verifyPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "remember", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.UserProjects)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GETIndex",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					23: []string{ 
						"projectList",
					},
				},
			},
			&revel.MethodType{
				Name: "POSTIndex",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "projectName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "projectDescription", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					34: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Hotels)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					40: []string{ 
						"bookings",
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "search", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "size", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "page", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					61: []string{ 
						"hotels",
						"search",
						"size",
						"page",
						"nextPage",
					},
				},
			},
			&revel.MethodType{
				Name: "Show",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					92: []string{ 
						"title",
						"hotel",
					},
				},
			},
			&revel.MethodType{
				Name: "Settings",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					96: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "SaveSettings",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "verifyPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ConfirmBooking",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "booking", Type: reflect.TypeOf((*models.Booking)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					147: []string{ 
						"title",
						"hotel",
						"booking",
					},
				},
			},
			&revel.MethodType{
				Name: "CancelBooking",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Book",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					166: []string{ 
						"title",
						"hotel",
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/akalenda/GolangRevelRBAC/app/controllers.Hotels.SaveSettings": { 
			101: "verifyPassword",
			103: "verifyPassword",
		},
		"github.com/akalenda/GolangRevelRBAC/app/models.(*Hotel).Validate": { 
			22: "hotel.Name",
			24: "hotel.Address",
			29: "hotel.City",
			35: "hotel.State",
			41: "hotel.Zip",
			47: "hotel.Country",
		},
		"github.com/akalenda/GolangRevelRBAC/app/models.(*User).Validate": { 
			88: "u.Username",
			93: "u.Name",
		},
		"github.com/akalenda/GolangRevelRBAC/app/models.Booking.Validate": { 
			44: "booking.User",
			45: "booking.Hotel",
			46: "booking.CheckInDate",
			47: "booking.CheckOutDate",
			49: "booking.CardNumber",
			56: "booking.NameOnCard",
		},
		"github.com/akalenda/GolangRevelRBAC/app/models.ValidatePassword": { 
			101: "password",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.ApplicationTest)(nil),
	}

	revel.Run(*port)
}
