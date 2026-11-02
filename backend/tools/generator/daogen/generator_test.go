//go:build unit

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/require"
)

// Compile and execute a generated operation to verify hook ordering, isolation,
// nil handling, and suppression of results when the operation or a hook fails.
func TestGeneratedHookList(t *testing.T) {
	dir := t.TempDir()
	cfg := &DaoCfg{InterfaceName: "Dao"}
	f := jen.NewFile("fixture")
	params := jen.Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("value").Int(),
		jen.Id("operationErr").Error(),
		jen.Id("callOpts").Op("...").Id("DaoRunOption"),
	)
	generateHookOptions(f, cfg, "Run", jen.Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("result").Int()), jen.Int())
	f.Var().Id("operationFinished").Bool()
	f.Func().Id("Run").Add(params).Params(jen.Int(), jen.Error()).Block(methodBody(cfg, "Run", []jen.Code{
		jen.Id("ctx"), jen.Id("result"),
	}, jen.Return(jen.Lit(0), jen.Id("err")),
		jen.Id("result").Op(":=").Id("value").Op("+").Len(jen.Id("opts")),
		jen.Id("operationFinished").Op("=").True(),
		jen.Id("err").Op(":=").Id("operationErr"),
		jen.Return(jen.Id("result"), jen.Id("err")),
	)...)
	require.NoError(t, f.Save(filepath.Join(dir, "operation.go")))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module fixture\n\ngo 1.26\n"), 0o600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "operation_test.go"), []byte(`package fixture

import (
    "context"
    "errors"
    "reflect"
    "testing"
)

func TestHooks(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    for _, operationErr := range []error{nil, errors.New("database error")} {
        operationFinished = false
        var calls []int
        hook := func(index int) func(context.Context, int) error {
            return func(hookCtx context.Context, result int) error {
                if !operationFinished || hookCtx != ctx || result != 44 {
                    t.Fatal("hook ran before operation or received incorrect arguments")
                }
                calls = append(calls, index)
                return nil
            }
        }
        result, err := Run(ctx, 42, operationErr,
            WithDaoRunHooks(hook(1), nil, hook(2)),
            WithDaoRunHook(hook(3)),
            WithDaoRunDatabaseOptions(1, 2),
        )
        var expectedCalls []int
        if operationErr == nil {
            expectedCalls = []int{1, 2, 3}
        }
        expectedResult := 44
        if operationErr != nil {
            expectedResult = 0
        }
        if result != expectedResult || err != operationErr || !reflect.DeepEqual(calls, expectedCalls) {
            t.Fatalf("result=%d err=%v calls=%v", result, err, calls)
        }
        result, err = Run(ctx, 42, nil)
        if result != 42 || err != nil || len(calls) != len(expectedCalls) {
            t.Fatal("hooks or database options leaked into another invocation")
        }
    }
    hookErr := errors.New("hook failed")
    for _, operationErr := range []error{nil, errors.New("database failed")} {
        called := false
        firstCalled := false
        result, err := Run(ctx, 42, operationErr, WithDaoRunHooks(
            func(context.Context, int) error { firstCalled = true; return hookErr },
            func(context.Context, int) error { called = true; return nil },
        ))
        expectedErr := hookErr
        if operationErr != nil {
            expectedErr = operationErr
        }
        if result != 0 || err != expectedErr || called || firstCalled != (operationErr == nil) {
            t.Fatalf("failure returned a result, lost the error, or ran subsequent hooks: %d %v %v", result, err, called)
        }
    }

}
`), 0o600))
	cmd := exec.CommandContext(t.Context(), "go", "test", "./...")
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOWORK=off")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "%s", output)
}

func TestLegacyMethodParameters(t *testing.T) {
	cfg := &DaoCfg{DisableV2: true}
	params := []jen.Code{jen.Id("ctx").Qual("context", "Context"), jen.Id("opts").Op("...").Int()}
	require.Contains(t, jen.Func().Id("Legacy").Add(methodParameters(cfg, params, "UnusedOption")).Block().GoString(),
		"func Legacy(ctx context.Context, opts ...int)")
	f := jen.NewFile("fixture")
	generateHookOptions(f, cfg, "Run", jen.Params(), jen.Int())
	require.NotContains(t, f.GoString(), "Hook")
}
