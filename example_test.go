package xmldom_test

import (
	"fmt"

	"github.com/subchen/go-xmldom"
)

const (
	ExampleXml = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE junit SYSTEM "junit-result.dtd">
<testsuites>
	<testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom">
		<properties>
			<property name="go.version">go1.8.1</property>
		</properties>
		<testcase classname="go-xmldom" name="ExampleParseXML" time="0.004"></testcase>
		<testcase classname="go-xmldom" name="ExampleParse" time="0.005"></testcase>
	</testsuite>
</testsuites>
`
)

func ExampleParseXML() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	fmt.Printf("name = %v\n", node.Name)
	fmt.Printf("attributes.len = %v\n", len(node.Attributes))
	fmt.Printf("children.len = %v\n", len(node.Children))
	fmt.Printf("root = %v", node == node.Root())
	// Output:
	// name = testsuites
	// attributes.len = 0
	// children.len = 1
	// root = true
}

func ExampleNode_GetAttribute() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	attr := node.FirstChild().GetAttribute("name")
	fmt.Printf("%v = %v\n", attr.Name, attr.Value)
	// Output:
	// name = github.com/subchen/go-xmldom
}

func ExampleNode_GetChildren() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	children := node.FirstChild().GetChildren("testcase")
	for _, c := range children {
		fmt.Printf("%v: name = %v\n", c.Name, c.GetAttributeValue("name"))
	}
	// Output:
	// testcase: name = ExampleParseXML
	// testcase: name = ExampleParse
}

func ExampleNode_Query() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	// xpath expr: https://github.com/antchfx/xpath

	// find all children
	fmt.Printf("children = %v\n", len(node.Query("//*")))

	// find node matched tag name
	nodeList := node.Query("//testcase")
	for _, c := range nodeList {
		fmt.Printf("%v: name = %v\n", c.Name, c.GetAttributeValue("name"))
	}
	// Output:
	// children = 5
	// testcase: name = ExampleParseXML
	// testcase: name = ExampleParse
}

func ExampleNode_QueryOne() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	// xpath expr: https://github.com/antchfx/xpath

	// find node matched attr name
	c := node.QueryOne("//testcase[@name='ExampleParseXML']")
	fmt.Printf("%v: name = %v\n", c.Name, c.GetAttributeValue("name"))
	// Output:
	// testcase: name = ExampleParseXML
}

func ExampleDocument_XML() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleXml))
	fmt.Println(doc.XML())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?><!DOCTYPE junit SYSTEM "junit-result.dtd"><testsuites><testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom"><properties><property name="go.version">go1.8.1</property></properties><testcase classname="go-xmldom" name="ExampleParseXML" time="0.004" /><testcase classname="go-xmldom" name="ExampleParse" time="0.005" /></testsuite></testsuites>
}

func ExampleDocument_XMLPretty() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleXml))
	fmt.Println(doc.XMLPretty())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <!DOCTYPE junit SYSTEM "junit-result.dtd">
	// <testsuites>
	//   <testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/go-xmldom">
	//     <properties>
	//       <property name="go.version">go1.8.1</property>
	//     </properties>
	//     <testcase classname="go-xmldom" name="ExampleParseXML" time="0.004" />
	//     <testcase classname="go-xmldom" name="ExampleParse" time="0.005" />
	//   </testsuite>
	// </testsuites>
}

func ExampleNewDocument() {
	doc := xmldom.NewDocument("testsuites")

	testsuiteNode := xmldom.NewNode("testsuite").SetAttributeValue("name", "github.com/subchen/go-xmldom")
	doc.Root.AppendChild(testsuiteNode)

	caseNode1 := xmldom.NewTextNode("testcase", "PASS")
	caseNode2 := xmldom.NewTextNode("testcase", "PASS")
	testsuiteNode.AppendChild(caseNode1).AppendChild(caseNode2)

	fmt.Println(doc.XML())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?><testsuites><testsuite name="github.com/subchen/go-xmldom"><testcase>PASS</testcase><testcase>PASS</testcase></testsuite></testsuites>
}
