package workflows

import (
  "backend/app/activitys"
  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/require"
  "testing"

  "go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
  // Set up the test suite and testing execution environment
  testSuite := &testsuite.WorkflowTestSuite{}
  env := testSuite.NewTestWorkflowEnvironment()

  // Mock activity implementation
  env.OnActivity(activitys.ComposeGreeting, mock.Anything, "World").Return("Hello World!", nil)

  env.ExecuteWorkflow(GreetingWorkflow, "World")
  require.True(t, env.IsWorkflowCompleted())
  require.NoError(t, env.GetWorkflowError())

  var greeting string
  require.NoError(t, env.GetWorkflowResult(&greeting))
  require.Equal(t, "Hello World!", greeting)

}
