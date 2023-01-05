package lang

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/j178/leetgo/config"
	"github.com/j178/leetgo/leetcode"
)

const (
	testutilsModPath = "github.com/j178/leetgo/testutils/go"
	testFileTemplate = `// Code generated by https://github.com/j178/leetgo

package main

import (
    "testing"

    . "%s"
)

var testcases = ` + "`" + `
%s
` + "`" + `

func Test_%s(t *testing.T) {
    targetCaseNum := 0
    // targetCaseNum := -1
    if err := %s(t, %s, testcases, targetCaseNum); err != nil {
        t.Fatal(err)
    }
}
`
)

type golang struct {
	baseLang
}

func addNamedReturn(code string, q *leetcode.QuestionData) string {
	lines := strings.Split(code, "\n")
	var newLines []string
	skipNext := 0
	for _, line := range lines {
		if skipNext > 0 {
			skipNext--
			continue
		}
		if strings.HasPrefix(line, "func ") {
			rightBrace := strings.LastIndex(line, ")")
			returnType := strings.TrimSpace(line[rightBrace+1 : strings.LastIndex(line, "{")])
			if returnType != "" {
				if returnType == "bool" || returnType == "string" {
					newLines = append(newLines, line)
				} else if q.MetaData.SystemDesign && strings.Contains(line, "func Constructor") {
					newLines = append(newLines, line)
					newLines = append(newLines, "\n\treturn "+returnType+"{}")
					skipNext = 1
				} else {
					newLines = append(newLines, line[:rightBrace+1]+" (ans "+returnType+") {")
					newLines = append(newLines, "\n\treturn")
					skipNext = 1
				}
			} else {
				newLines = append(newLines, line)
			}
		} else {
			newLines = append(newLines, line)
		}
	}
	return strings.Join(newLines, "\n")
}

func changeReceiverName(code string, q *leetcode.QuestionData) string {
	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "func (this *") {
			n := len("func (this *")
			prefix := strings.ToLower(line[n : n+1])
			lines[i] = strings.Replace(line, "this", prefix, 1)
		}
	}
	return strings.Join(lines, "\n")
}

func (g golang) Initialized(outDir string) (bool, error) {
	cmd := exec.Command("go", "list", "-m", "-json", testutilsModPath)
	cmd.Dir = outDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		if bytes.Contains(output, []byte("not a known dependency")) {
			return false, nil
		}
		return false, fmt.Errorf("go list failed: %w", err)
	}
	return true, nil
}

func (g golang) Init(outDir string) error {
	modPath := config.Get().Code.Go.GoModPath
	if modPath == "" {
		modPath = "leetcode-solutions"
		hclog.L().Warn("GoModPath path is not set, use default path", "mod_path", modPath)
	}
	var stderr bytes.Buffer
	cmd := exec.Command("go", "mod", "init", modPath)
	cmd.Dir = outDir
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)
	err := cmd.Run()
	if err != nil && !bytes.Contains(stderr.Bytes(), []byte("go.mod already exists")) {
		return err
	}

	cmd = exec.Command("go", "get", "-u", testutilsModPath)
	cmd.Dir = outDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}

func (g golang) RunTest(q *leetcode.QuestionData) error {
	// TODO run go test
	// Need go list -m json to get mod path
	return nil
}

func (g golang) generateTest(q *leetcode.QuestionData, testcases string) string {
	var funcName, testFuncName string
	if q.MetaData.SystemDesign {
		funcName = "Constructor"
		testFuncName = "RunClassTestsWithString"
	} else {
		funcName = q.MetaData.Name
		testFuncName = "RunTestsWithString"
	}
	content := fmt.Sprintf(testFileTemplate, testutilsModPath, testcases, funcName, testFuncName, funcName)
	return content
}

func (g golang) GeneratePaths(q *leetcode.QuestionData) (*GenerateResult, error) {
	filenameTmpl := getFilenameTemplate(g)
	baseFilename, err := q.GetFormattedFilename(g.slug, filenameTmpl)
	if err != nil {
		return nil, err
	}
	codeFile := filepath.Join(baseFilename, "solution.go")
	testFile := filepath.Join(baseFilename, "solution_test.go")

	files := []FileOutput{
		{
			Path: codeFile,
			Type: CodeFile,
		},
		{
			Path: testFile,
			Type: TestFile,
		},
	}

	return &GenerateResult{
		Generator: g,
		Files:     files,
	}, nil
}

func (g golang) Generate(q *leetcode.QuestionData) (*GenerateResult, error) {
	comment := g.generateComments(q)
	code := q.GetCodeSnippet(g.Slug())
	preCode := "package main\n\n"
	if needsDefinition(code) {
		preCode += fmt.Sprintf("import . \"%s\"\n\n", testutilsModPath)
	}
	code = g.generateCode(
		q,
		removeComments,
		addNamedReturn,
		changeReceiverName,
		addCodeMark(g),
		prepend(preCode),
	)
	codeContent := comment + "\n" + code + "\n"

	testcaseStr := g.generateTestCases(q)
	testContent := g.generateTest(q, testcaseStr)

	filenameTmpl := getFilenameTemplate(g)
	baseFilename, err := q.GetFormattedFilename(g.slug, filenameTmpl)
	if err != nil {
		return nil, err
	}
	codeFile := filepath.Join(baseFilename, "solution.go")
	testFile := filepath.Join(baseFilename, "solution_test.go")

	files := []FileOutput{
		{
			Path:    codeFile,
			Content: codeContent,
			Type:    CodeFile,
		},
		{
			Path:    testFile,
			Content: testContent,
			Type:    TestFile,
		},
	}

	return &GenerateResult{
		Generator: g,
		Files:     files,
	}, nil
}
