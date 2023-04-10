package traversal

// https://tinkerpop.apache.org/javadocs/3.6.2/core/org/apache/tinkerpop/gremlin/process/traversal/dsl/graph/GraphTraversal.html#times(int)

// Times Modifies a repeat(Traversal) to specify how many loops should occur before exiting.
// maxLoops - the number of loops to execute prior to exiting
// Signatures:
// times​(int)
func (g String) Times(maxLoops int) String {
	g.AddStep("times", maxLoops)

	return g
}
