package airbyte

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/harshithmullapudi/airbyte/logger"
	"github.com/harshithmullapudi/airbyte/models"
	"github.com/jedib0t/go-pretty/v6/table"
)

//Sources Print

func PrintSourcesTable(sources models.Sources) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Source Definition Id", "Source Id", "Name", "Source Name"})
	for index, s := range sources {
		t.AppendRow([]interface{}{index + 1, s.SourceDefinitionId, s.SourceId, s.Name, s.SourceName})
	}
	t.Render()
}

func PrintSourceTable(source models.Source) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Source Definition Id", "Source Id", "Name", "Source Name"})

	t.AppendRow([]interface{}{1, source.SourceDefinitionId, source.SourceId, source.Name, source.SourceName})

	t.Render()
}

func PrintSources(sources models.Sources) {

	b, err := json.MarshalIndent(sources, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

func PrintSource(source models.Source) {

	b, err := json.MarshalIndent(source, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

//Conenctions print

func PrintConnectionsTable(connections models.Connections) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Connection Id", "Name", "Source Id", "Source Name", "Destination Id", "Schedule", "Status"})
	for index, c := range connections {
		t.AppendRow([]interface{}{index + 1, c.ConnectionId, c.Name, c.SourceId, c.Source.Name, c.DestinationId, c.Schedule, c.Status})
	}
	t.Render()
}

func PrintConnectionTable(connection models.Connection) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Connection Id", "Name", "Source Id", "Source Name", "Destination Id", "Schedule", "Status"})
	t.AppendRow([]interface{}{1, connection.ConnectionId, connection.Name, connection.SourceId, connection.Source.Name, connection.DestinationId, connection.Schedule, connection.Status})
	t.Render()
}

func PrintConnections(connections models.Connections) {

	b, err := json.MarshalIndent(connections, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}

func PrintConnection(connection models.Connection) {

	b, err := json.MarshalIndent(connection, "", "  ")

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(string(b))
}
