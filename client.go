package hdfs

import (
	"code.google.com/p/goprotobuf/proto"
	hdfs "github.com/colinmarc/hdfs/protocol/hadoop_hdfs"
	"github.com/colinmarc/hdfs/rpc"
	"io/ioutil"
	"os"
	"os/user"
)

// A Client represents a connection to an HDFS cluster
type Client struct {
	namenode *rpc.NamenodeConnection
}

// New returns a connected Client, or an error if it can't connect
func New(address string) (*Client, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	namenode, err := rpc.NewNamenodeConnection(address, currentUser.Username)
	if err != nil {
		return nil, err
	}

	return &Client{namenode: namenode}, nil
}

// Chmod changes the mode of the named file to mode.
func (c *Client) Chmod(name string, mode os.FileMode) error {
	return nil
}

// Chown changes the numeric uid and gid of the named file.
func (c *Client) Chown(name string, uid, gid int) error {
	return nil
}

// Mkdir creates a new directory with the specified name and permission bits.
func (c *Client) Mkdir(name string, perm os.FileMode) error {
	return nil
}

// MkdirAll creates a directory named path, along with any necessary parents,
// and returns nil, or else returns an error. The permission bits perm are used
// for all directories that MkdirAll creates. If path is already a directory,
// MkdirAll does nothing and returns nil.
func (c *Client) MkdirAll(path string, perm os.FileMode) error {
	return nil
}

// Remove removes the named file or directory.
func (c *Client) Remove(name string) error {
	return nil
}

// Rename renames (moves) a file.
func (c *Client) Rename(oldpath, newpath string) error {
	return nil
}

// Stat returns an os.FileInfo describing the named file.
func (c *Client) Stat(name string) (fi os.FileInfo, err error) {
	req := &hdfs.GetFileInfoRequestProto{Src: proto.String(name)}
	resp := &hdfs.GetFileInfoResponseProto{}

	err = c.namenode.Execute("getFileInfo", req, resp)
	if err != nil {
		return nil, err
	}

	return &FileInfo{name, resp.GetFs()}, nil
}

// ReadDir reads the directory named by dirname and returns a list of sorted
// directory entries.
func (c *Client) ReadDir(dirname string) ([]os.FileInfo, error) {
	return []os.FileInfo{}, nil
}

// ReadFile reads the file named by filename and returns the contents.
func (c *Client) ReadFile(filename string) ([]byte, error) {
	f, err := c.Open(filename)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

// CopyToLocal copies the HDFS file specified by src to the local file at dst.
// If dst already exists, it will be overwritten.
func (c *Client) CopyToLocal(src string, dst string) error {
	return nil
}