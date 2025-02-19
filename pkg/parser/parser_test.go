package parser_test

// import (
// 	"fmt"
// 	"some_app/pkg/parser"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// //TODO: написать тесты для парсера
// func TestParserHH(t *testing.T) {
// 	cases := []struct{
// 		In  []string
// 		Want bool
// 		}{
// 			{
// 			In : []string{"php", "1", "100"},
// 			Want : false,
// 			},
// 			{
// 			In : []string{"php", "100", "100"},
// 			Want : true,
// 			},
// 		}

// 	for _, tc := range cases {
// 		testTitle := fmt.Sprintf("lang %s; cont page %s; on page %s", tc.In[0], tc.In[1], tc.In[2])

// 		t.Run(testTitle, func(t *testing.T) {
// 			someParser := parser.NewHTTPParseClient(tc.In[0])
// 			parsResult, err := someParser.RecvData(tc.In[0], tc.In[1], tc.In[2])
// 			require.Nil(t, parsResult)
// 			require.Nil(t, err)
// 		})
// 	}
// }