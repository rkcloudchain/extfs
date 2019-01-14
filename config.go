package extfs

// Config epresents the configurable options for a filesystem.
type Config struct {
	// User specifies which HDFS user the client will act as. HDFS only
	User string

	// Addresses specifies the namenode(s) to connect to. HDFS only
	Addresses []string
}
