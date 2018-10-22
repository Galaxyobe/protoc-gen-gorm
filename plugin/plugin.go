package plugin

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"
	"io/ioutil"
	"strings"
	"log"
	"os"
	"bufio"
	"io"
	"bytes"
	"time"
	"os/exec"
	"strconv"
	"syscall"
	"html/template"
	"github.com/gogo/protobuf/proto"
	"github.com/galaxyobe/protoc-gen-gorm/proto"
	"errors"
)

const (
	contextPkg = "context"
	GORMPkg    = "github.com/jinzhu/gorm"
	TimePkg    = "time"
)

type generateData struct {
	Package     string
	MessageName string
	ContextPkg  string
	GORMPkg     string
	TimePkg     string
	// field
	PrimaryKey string
	CreateAt   string
	UpdateAt   string
	DeleteAt   string
}

type Plugin struct {
	*generator.Generator
	generator.PluginImports
	UseGogoImport bool
	GeneratePath  string
}

func NewPlugin(useGogoImport bool, generatePath string) *Plugin {
	return &Plugin{
		UseGogoImport: useGogoImport,
		GeneratePath:  generatePath,
	}
}

func (p *Plugin) Name() string {
	return "gorm"
}

func (p *Plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *Plugin) Generate(file *generator.FileDescriptor) {
	if len(file.Messages()) == 0 {
		return
	}

	if !p.UseGogoImport {
		vanity.TurnOffGogoImport(file.FileDescriptorProto)
	}

	p.PluginImports = generator.NewPluginImports(p.Generator)

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		if err := p.generateGORMFunc(file, msg); err != nil {
			panic(err)
		}
	}
}

func (p *Plugin) InjectIgnoreFork() error {

	lp, err := exec.LookPath(os.Args[0])
	if err != nil {
		return err
	}

	argv := []string{lp, "-inject", "-inject-path=" + p.GeneratePath, "-ppid=" + strconv.Itoa(os.Getpid())}

	_, err = os.StartProcess(lp, argv, &os.ProcAttr{})
	if err != nil {
		return err
	}

	return nil
}

func (p *Plugin) InjectIgnore(ppid int) error {

	for {
		if process, err := os.FindProcess(ppid); err != nil {
			break
		} else {
			if err = process.Signal(syscall.Signal(0)); err != nil {
				break
			}
		}
		time.Sleep(1 * time.Second)
	}

	files, err := ioutil.ReadDir(p.GeneratePath)
	if err != nil {
		return err
	}

	var nn []string
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".pb.go") || file.IsDir() {
			continue
		}
		nn = append(nn, file.Name())
		log.Println(file.Name())
		filePath := p.GeneratePath + "/" + file.Name()
		f, err := os.OpenFile(filePath, os.O_RDWR, 0660)
		if err != nil {
			log.Println("ReadFile", file, err)
			continue
		}

		reader := bufio.NewReader(f)

		buffer := bytes.NewBuffer(make([]byte, 0))

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				f.Close()
				return err
			}

			if bytes.Contains(line, []byte(`json:"-"`)) && !bytes.Contains(line, []byte(`gorm:"-"`)) {
				log.Println(string(line))
				data := bytes.Replace(line, []byte(`json:"-"`), []byte(`json:"-" gorm:"-"`), -1)
				log.Println(string(data))
				buffer.Write(data)
				buffer.WriteString("\n")
			} else {
				buffer.Write(line)
				buffer.WriteString("\n")
			}
		}
		f.WriteAt(buffer.Bytes(), 0)
		f.Close()
	}

	return nil
}

const gormControllerTemplate = `
// new {{.MessageName}} GORM controller with gorm.DB
func (m *{{.MessageName}}) GORMController(db *{{.GORMPkg}}.DB) *{{.MessageName}}GORMController {
	return &{{.MessageName}}GORMController{
		DB: db,
		m:  m,
	}
}

// new {{.MessageName}} GORM controller with gorm.DB
func New{{.MessageName}}GORMController(db *{{.GORMPkg}}.DB) *{{.MessageName}}GORMController {
	return &{{.MessageName}}GORMController{DB: db, m: new({{.MessageName}})}
}

type {{.MessageName}}GORMController struct {
	DB *{{.GORMPkg}}.DB
	m  *{{.MessageName}}
}

func (g *{{.MessageName}}GORMController) M(m *{{.MessageName}}) {
	g.m = m
}

func (g *{{.MessageName}}GORMController) M() *{{.MessageName}} {
	return g.m
}

func (g *{{.MessageName}}GORMController) AutoMigrate() {
	g.DB.AutoMigrate(g.m)
}

func (g *{{.MessageName}}GORMController) Begin() error {
	if db := g.DB.Begin(); db.Error != nil {
		return db.Error
	} else {
		g.DB = db
	}

	return nil
}

func (g *{{.MessageName}}GORMController) Rollback() *{{.GORMPkg}}.DB {
	return g.DB.Rollback()
}

func (g *{{.MessageName}}GORMController) Commit() *{{.GORMPkg}}.DB {
	return g.DB.Commit()
}

func (g *{{.MessageName}}GORMController) Create() *{{.GORMPkg}}.DB {
{{- if or .CreateAt .UpdateAt }}
	now := {{.TimePkg}}.Now().Unix()
{{- end }}
{{- if .CreateAt }}
	g.m.{{.CreateAt}} = now
{{- end }}
{{- if .UpdateAt }}
	g.m.{{.UpdateAt}} = now
{{- end }}

	return g.DB.Create(g.m)
}

{{ if .PrimaryKey }}
func (g *{{.MessageName}}GORMController) Delete() *{{.GORMPkg}}.DB {
	if g.m.{{.PrimaryKey}} == 0 {
		g.DB.Error = errors.New("the value of {{.PrimaryKey}} is not expected to be 0")
		return g.DB
	}

	return g.DB.Delete(g.m)
}
{{- end }}

{{ if .DeleteAt }}
func (g *{{.MessageName}}GORMController) SoftDelete() *{{.GORMPkg}}.DB {
	g.m.{{.DeleteAt}} = {{.TimePkg}}.Now().Unix()

	return g.DB.Model(g.m).Select("{{.DeleteAt}}").Updates(g.m)
}
{{- end }}

func (g *{{.MessageName}}GORMController) Update() *{{.GORMPkg}}.DB {
{{- if .UpdateAt }}
	g.m.{{.UpdateAt}} = {{.TimePkg}}.Now().Unix()
{{ end }}
	return g.DB.Model(g.m).Omit({{ if .PrimaryKey }} "{{.PrimaryKey}}" {{ end }}{{ if .CreateAt }} ,"{{.CreateAt}}" {{ end }}{{ if .DeleteAt }} ,"{{.DeleteAt}}" {{ end }}).Updates(g.m)
}

func (g *{{.MessageName}}GORMController) First() (*{{.MessageName}}, error) {
	db := g.DB.First(g.m)

	return g.m, db.Error
}

// when where is empty string will be ignored
// when order is empty is null: nil,-1,"" will be ignored
// when offset and limit is null: nil,-1,"" will be ignored
func (g *{{.MessageName}}GORMController) Find(where, limit, offset, order interface{}) ([]*{{.MessageName}}, error) {
	var array []*{{.MessageName}}

	db := g.DB.Where(where).Order(order).Limit(limit).Offset(offset).Find(&array)

	return array, db.Error
}

// when where is empty string will be ignored
// when order is empty is null: nil,-1,"" will be ignored
// when offset and limit is null: nil,-1,"" will be ignored
func (g *{{.MessageName}}GORMController) Count(where, limit, offset, order interface{}) (int64, error) {
	var count int64 = 0

	db := g.DB.Model(g.m).Where(where).Order(order).Limit(limit).Offset(offset).Count(&count)

	return count, db.Error
}
`

func (p *Plugin) generateGORMFunc(file *generator.FileDescriptor, message *generator.Descriptor) error {

	// enable gorm
	if proto.GetBoolExtension(message.Options, gorm.E_Enabled, false) {
		// generateData
		data := &generateData{
			ContextPkg:  p.NewImport(contextPkg).Use(),
			GORMPkg:     p.NewImport(GORMPkg).Use(),
			MessageName: generator.CamelCaseSlice(message.TypeName()),
		}

		// range field
		for _, field := range message.Field {

			if proto.GetBoolExtension(field.Options, gorm.E_PrimaryKey, false) {
				data.PrimaryKey = generator.CamelCase(*field.Name)
			} else if proto.GetBoolExtension(field.Options, gorm.E_CreateAt, false) {
				data.CreateAt = generator.CamelCase(*field.Name)
			} else if proto.GetBoolExtension(field.Options, gorm.E_UpdateAt, false) {
				data.UpdateAt = generator.CamelCase(*field.Name)
			} else if proto.GetBoolExtension(field.Options, gorm.E_DeleteAt, false) {
				data.DeleteAt = generator.CamelCase(*field.Name)
			}

			if field.TypeName != nil {
				// use external proto
				p.Generator.RecordTypeUse(*field.TypeName)
			}
		}

		// find gorm special field
		if data.PrimaryKey == "" || data.CreateAt == "" || data.UpdateAt == "" || data.DeleteAt == "" {
			for _, field := range message.Field {
				if data.PrimaryKey == "" && strings.Contains(field.Options.String(), `gorm:"primary_key"`) {
					data.PrimaryKey = generator.CamelCase(*field.Name)
				} else if data.CreateAt == "" && generator.CamelCase(*field.Name) == "CreateAt" {
					data.CreateAt = generator.CamelCase(*field.Name)
				} else if data.UpdateAt == "" && generator.CamelCase(*field.Name) == "UpdateAt" {
					data.UpdateAt = generator.CamelCase(*field.Name)
				} else if data.DeleteAt == "" && generator.CamelCase(*field.Name) == "DeleteAt" {
					data.DeleteAt = generator.CamelCase(*field.Name)
				}
			}
		}

		if data.PrimaryKey == "" {
			return errors.New("no primary key")
		}

		if data.CreateAt != "" || data.UpdateAt != "" || data.DeleteAt != "" {
			data.TimePkg = p.NewImport(TimePkg).Use()
		}

		tmpl, err := template.New("GORMController").Parse(gormControllerTemplate)
		if err != nil {
			return err
		}
		return tmpl.Execute(p.Buffer, data)
	}

	return nil
}
