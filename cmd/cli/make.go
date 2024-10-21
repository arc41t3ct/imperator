package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ettle/strcase"
	"github.com/fatih/color"
	pl "github.com/gertd/go-pluralize"
)

func makeCommand(arg2, arg3 string) error {
	switch arg2 {
	case "key":
		randString := imp.CreateRadomString(32)
		color.Green("32 character enncryption key: %s", randString)
	case "migration":
		if arg3 == "" {
			return errors.New("you must specify a name for the migration")
		}
		if err := makeMigration(arg3); err != nil {
			return errors.New(fmt.Sprintf("failed to create migration: %s", err))
		}
	case "auth":
		if err := implementAuth(); err != nil {
			return errors.New(fmt.Sprintf("failed to implement auth: %s", err))
		}
	case "handler":
		if arg3 == "" {
			return errors.New("you must specify a name for the handler")
		}
		if err := makeHandler(arg3); err != nil {
			return errors.New(fmt.Sprintf("failed to create handler: %s", err))
		}
	case "model":
		if arg3 == "" {
			return errors.New("you must specify a name for the model")
		}
		if err := makeModel(arg3); err != nil {
			return errors.New(fmt.Sprintf("failed to create model: %s", err))
		}
	case "middleware":
		if arg3 == "" {
			return errors.New("you must specify a name for the middleware")
		}
		if err := makeMiddleware(arg3); err != nil {
			return errors.New(fmt.Sprintf("failed to create middleware: %s", err))
		}
	case "session":
		if err := implementSession(); err != nil {
			return errors.New(fmt.Sprintf("failed to implement session: %s", err))
		}
	case "mail":
		if arg3 == "" {
			return errors.New("you must specify a name for the mail template")
		}
		if err := makeMail(arg3); err != nil {
			return errors.New(fmt.Sprintf("failed to create mail: %s", err))
		}
	}
	return nil
}

func makeMail(arg3 string) error {
	htmlMail := imp.RootPath + "/mail/" + strings.ToLower(arg3) + ".html.tmpl"
	plainMail := imp.RootPath + "/mail/" + strings.ToLower(arg3) + ".plain.tmpl"
	if err := copyFileFromTemplate(
		"templates/mail/mail.html.tmpl", htmlMail); err != nil {
		return errors.New(fmt.Sprintf("failed to copy html mail template: %s", err))
	}
	if err := copyFileFromTemplate(
		"templates/mail/mail.plain.tmpl", plainMail); err != nil {
		return errors.New(fmt.Sprintf("failed to copy text mail template: %s", err))
	}

	color.Green("[+] Mail has been created...")
	color.Yellow(" - added %s mail html and plain templates to /mail", arg3)
	fmt.Println("")
	color.Green("[!] Next manual steps...")
	color.Blue(" - remember edit mail file:")
	color.Blue("\t%s", htmlMail)
	color.Blue(" - remember edit mail file:")
	color.Blue("\t%s", plainMail)
	return nil
}

func makeMigration(arg3 string) error {
	dbType := imp.DB.DatabaseType
	fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)
	upFile := imp.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := imp.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	if err := copyFileFromTemplate(
		"templates/migrations/migration."+dbType+".up.sql", upFile); err != nil {
		return errors.New(fmt.Sprintf("failed to copy up template: %s", err))
	}

	if err := copyFileFromTemplate(
		"templates/migrations/migration."+dbType+".down.sql", downFile); err != nil {
		return errors.New(fmt.Sprintf("failed to copy down template: %s", err))
	}

	color.Green("[+] Migration has been created...")
	color.Yellow(" - added up migration: %s to /migrations", fileName)
	color.Yellow(" - added down migration: %s to /migrations", fileName)
	fmt.Println("")
	color.Green("[!] Next manual steps...")
	color.Blue(" - remember edit up migration file:")
	color.Blue("\t%s", upFile)
	color.Blue(" - remember edit down migration file:")
	color.Blue("\t%s", downFile)
	return nil
}

func makeHandler(arg3 string) error {
	fileName := imp.RootPath + "/handlers/" + strcase.ToSnake(arg3) + ".go"
	data, err := templateFS.ReadFile("templates/handlers/handler.go.txt")
	if err != nil {
		return errors.New(fmt.Sprintf("failed to find the handler's template: %s", err))
	}

	handlerName := strcase.ToGoPascal(arg3)
	handlerViewName := strcase.ToSnake(arg3)
	handler := string(data)
	handler = strings.ReplaceAll(handler, "{{HANDLER_NAME}}", handlerName)
	handler = strings.ReplaceAll(handler, "{{HANDLER_VIEW_NAME}}", handlerViewName)

	err = copyDataToFile([]byte(handler), fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to write the file: %s", err))
	}

	color.Green("[+] Handler has been created...")
	color.Yellow(" - added handler: %s to /handlers", handlerName)
	fmt.Println("")
	color.Green("[!] Next manual steps...")
	color.Blue(" - remember to edit %s handler in:", handlerName)
	color.Blue("\t%s", fileName)
	color.Blue(" - remember to add %s handler to /routes.go", handlerName)
	return nil
}

func makeModel(arg3 string) error {
	plural := pl.NewClient()
	if plural.IsPlural(arg3) {
		arg3 = plural.Singular(arg3)
	}
	fileName := imp.RootPath + "/models/" + strcase.ToSnake(arg3) + ".go"
	data, err := templateFS.ReadFile("templates/models/model.go.txt")
	if err != nil {
		return errors.New(fmt.Sprintf("failed to find the model's template: %s", err))
	}

	modelName := strcase.ToGoPascal(arg3)
	modelTableName := strcase.ToSnake(arg3)
	modelTableName = plural.Plural(modelTableName)
	model := string(data)
	model = strings.ReplaceAll(model, "{{MODEL_NAME}}", modelName)
	model = strings.ReplaceAll(model, "{{MODEL_TABLE_NAME}}", modelTableName)

	err = copyDataToFile([]byte(model), fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to write the file: %s", err))
	}

	color.Green("[+] Model has been created...")
	color.Yellow(" - added model: %s to /models", modelName)
	fmt.Println("")
	color.Green("[!] Next manual steps...")
	color.Blue(" - remember to edit %s model in:", modelName)
	color.Blue("\t%s", fileName)
	color.Blue(" - remember to add %s model to /models/models.go", modelName)
	return nil
}

func makeMiddleware(arg3 string) error {
	fileName := imp.RootPath + "/middleware/" + strcase.ToSnake(arg3) + ".go"
	data, err := templateFS.ReadFile("templates/middleware/middleware.go.txt")
	if err != nil {
		return errors.New(fmt.Sprintf("failed to find the middleware's template: %s", err))
	}

	middlewareName := strcase.ToGoPascal(arg3)
	middleware := string(data)
	middleware = strings.ReplaceAll(middleware, "{{MIDDLEWARE_NAME}}", middlewareName)

	err = copyDataToFile([]byte(middleware), fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to write the file: %s", err))
	}

	color.Green("[+] Middleware has been created...")
	color.Yellow(" - added middleware: %s to /middleware", middlewareName)
	fmt.Println("")
	color.Green("[!] Next manual steps...")
	color.Blue(" - remember to edit %s middleware in:", middlewareName)
	color.Blue("\t%s", fileName)
	color.Blue(" - remember to add %s middleware to /routes.go ", middlewareName)
	return nil
}
