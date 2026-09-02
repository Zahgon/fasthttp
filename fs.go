package fasthttp

import (
	"errors"
	"io"
	"io/fs"
	"sync"
	"time"
)

func ServeFileBytesUncompressed(ctx *RequestCtx, path []byte) { _ = "STUB: not implemented"; return }

func ServeFileUncompressed(ctx *RequestCtx, path string) { _ = "STUB: not implemented"; return }

func ServeFileBytes(ctx *RequestCtx, path []byte) { _ = "STUB: not implemented"; return }

func ServeFile(ctx *RequestCtx, path string) { _ = "STUB: not implemented"; return }

func ServeFileLiteral(ctx *RequestCtx, path string) { _ = "STUB: not implemented"; return }

var (
	rootFSOnce sync.Once
	rootFS     = &FS{
		Root:               "",
		AllowEmptyRoot:     true,
		GenerateIndexPages: true,
		Compress:           true,
		CompressBrotli:     true,
		CompressZstd:       true,
		AcceptByteRange:    true,
	}
	rootFSHandler RequestHandler
)

func ServeFS(ctx *RequestCtx, filesystem fs.FS, path string) { _ = "STUB: not implemented"; return }

func ServeFSLiteral(ctx *RequestCtx, filesystem fs.FS, path string) {
	_ = "STUB: not implemented"
	return
}

func serveFS(ctx *RequestCtx, filesystem fs.FS, path string, literal bool) {
	_ = "STUB: not implemented"
	return
}

func normalizeServeFilePath(ctx *RequestCtx, path string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

type PathRewriteFunc func(ctx *RequestCtx) []byte

func NewVHostPathRewriter(slashesCount int) PathRewriteFunc {
	_ = "STUB: not implemented"
	return *new(PathRewriteFunc)
}

var strInvalidHost = []byte("invalid-host")

func NewPathSlashesStripper(slashesCount int) PathRewriteFunc {
	_ = "STUB: not implemented"
	return *new(PathRewriteFunc)
}

func NewPathPrefixStripper(prefixSize int) PathRewriteFunc {
	_ = "STUB: not implemented"
	return *new(PathRewriteFunc)
}

type FS struct {
	noCopy noCopy

	FS fs.FS

	PathRewrite PathRewriteFunc

	PathNotFound RequestHandler

	CompressedFileSuffixes map[string]string

	CleanStop chan struct{}

	h RequestHandler

	Root string

	CompressRoot string

	CompressedFileSuffix string

	IndexNames []string

	CacheDuration time.Duration

	once sync.Once

	AllowEmptyRoot bool

	CompressBrotli bool

	CompressZstd bool

	GenerateIndexPages bool

	Compress bool

	AcceptByteRange bool

	SkipCache bool
}

const FSCompressedFileSuffix = ".fasthttp.gz"

var FSCompressedFileSuffixes = map[string]string{
	"gzip": ".fasthttp.gz",
	"br":   ".fasthttp.br",
	"zstd": ".fasthttp.zst",
}

const FSHandlerCacheDuration = 10 * time.Second

func FSHandler(root string, stripSlashes int) RequestHandler {
	_ = "STUB: not implemented"
	return *new(RequestHandler)
}

func (fs *FS) NewRequestHandler() RequestHandler {
	_ = "STUB: not implemented"
	return *new(RequestHandler)
}

func (fs *FS) normalizeRoot(root string) string { _ = "STUB: not implemented"; return "" }

func (fs *FS) initRequestHandler() { _ = "STUB: not implemented"; return }

type fsHandler struct {
	smallFileReaderPool sync.Pool
	filesystem          fs.FS

	cacheManager cacheManager

	pathRewrite            PathRewriteFunc
	pathNotFound           RequestHandler
	compressedFileSuffixes map[string]string

	root               string
	compressRoot       string
	indexNames         []string
	generateIndexPages bool
	compress           bool
	compressBrotli     bool
	compressZstd       bool
	acceptByteRange    bool
}

type fsFile struct {
	lastModified time.Time

	t               time.Time
	f               fs.File
	h               *fsHandler
	filename        string
	contentType     string
	dirIndex        []byte
	lastModifiedStr []byte

	bigFiles      []*bigFileReader
	contentLength int
	readersCount  int

	bigFilesLock sync.Mutex
	compressed   bool
}

func (ff *fsFile) NewReader() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (ff *fsFile) smallFileReader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

//nolint:forcetypeassert

const maxSmallFileSize = 2 * 4096

func (ff *fsFile) isBig() bool { _ = "STUB: not implemented"; return false }

func (ff *fsFile) bigFileReader() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (ff *fsFile) Release() { _ = "STUB: not implemented"; return }

func (ff *fsFile) decReadersCount() { _ = "STUB: not implemented"; return }

type bigFileReader struct {
	f  fs.File
	ff *fsFile
	r  io.Reader
	lr io.LimitedReader
}

func (r *bigFileReader) UpdateByteRange(startPos, endPos int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *bigFileReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *bigFileReader) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *bigFileReader) Close() error { _ = "STUB: not implemented"; return nil }

type fsSmallFileReader struct {
	ff       *fsFile
	startPos int
	endPos   int
}

func (r *fsSmallFileReader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *fsSmallFileReader) UpdateByteRange(startPos, endPos int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *fsSmallFileReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *fsSmallFileReader) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:forcetypeassert

type cacheManager interface {
	Lock()
	Unlock()
	Close()
	DecReadersCount(ff *fsFile)
	GetFileFromCache(cacheKind CacheKind, path []byte) (*fsFile, bool)
	SetFileToCache(cacheKind CacheKind, path []byte, ff *fsFile) *fsFile
}

var (
	_ cacheManager = (*inMemoryCacheManager)(nil)
	_ cacheManager = (*noopCacheManager)(nil)
)

type CacheKind uint8

const (
	defaultCacheKind CacheKind = iota
	brotliCacheKind
	gzipCacheKind
	zstdCacheKind
)

func newCacheManager(fs *FS) cacheManager { _ = "STUB: not implemented"; return *new(cacheManager) }

type noopCacheManager struct {
	cacheLock sync.Mutex
}

func (n *noopCacheManager) Lock() { _ = "STUB: not implemented"; return }

func (n *noopCacheManager) Unlock() { _ = "STUB: not implemented"; return }

func (*noopCacheManager) Close() { _ = "STUB: not implemented"; return }

func (n *noopCacheManager) DecReadersCount(ff *fsFile) { _ = "STUB: not implemented"; return }

func (*noopCacheManager) GetFileFromCache(cacheKind CacheKind, path []byte) (*fsFile, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (n *noopCacheManager) SetFileToCache(cacheKind CacheKind, path []byte, ff *fsFile) *fsFile {
	_ = "STUB: not implemented"
	return nil
}

type inMemoryCacheManager struct {
	cache         map[string]*fsFile
	cacheBrotli   map[string]*fsFile
	cacheGzip     map[string]*fsFile
	cacheZstd     map[string]*fsFile
	cacheDuration time.Duration
	cleanStop     chan struct{}
	cleanStopOnce sync.Once
	pendingFiles  []*fsFile
	closed        bool
	cacheLock     sync.Mutex
}

func (cm *inMemoryCacheManager) Lock() { _ = "STUB: not implemented"; return }

func (cm *inMemoryCacheManager) Unlock() { _ = "STUB: not implemented"; return }

func (cm *inMemoryCacheManager) Close() { _ = "STUB: not implemented"; return }

func (cm *inMemoryCacheManager) close() []*fsFile { _ = "STUB: not implemented"; return nil }

func (cm *inMemoryCacheManager) DecReadersCount(ff *fsFile) { _ = "STUB: not implemented"; return }

func (cm *inMemoryCacheManager) getFsCache(cacheKind CacheKind) map[string]*fsFile {
	_ = "STUB: not implemented"
	return nil
}

func (cm *inMemoryCacheManager) GetFileFromCache(cacheKind CacheKind, path []byte) (*fsFile, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (cm *inMemoryCacheManager) SetFileToCache(cacheKind CacheKind, path []byte, ff *fsFile) *fsFile {
	_ = "STUB: not implemented"
	return nil
}

func (cm *inMemoryCacheManager) handleCleanCache(cleanStop chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (cm *inMemoryCacheManager) cleanCache() []*fsFile { _ = "STUB: not implemented"; return nil }

func (cm *inMemoryCacheManager) cleanCacheNolock(cache map[string]*fsFile, filesToRelease []*fsFile) []*fsFile {
	_ = "STUB: not implemented"
	return nil
}

func (cm *inMemoryCacheManager) collectAllFilesToReleaseNolock(filesToRelease []*fsFile) []*fsFile {
	_ = "STUB: not implemented"
	return nil
}

func (cm *inMemoryCacheManager) collectCacheFilesToReleaseNolock(cache map[string]*fsFile, filesToRelease []*fsFile) []*fsFile {
	_ = "STUB: not implemented"
	return nil
}

func (cm *inMemoryCacheManager) addFileToReleaseNolock(filesToRelease []*fsFile, ff *fsFile) []*fsFile {
	_ = "STUB: not implemented"
	return nil
}

func (cm *inMemoryCacheManager) removePendingFileNolock(ff *fsFile) {
	_ = "STUB: not implemented"
	return
}

func (h *fsHandler) pathToFilePath(path []byte, hasTrailingSlash bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *fsHandler) filePathToCompressed(filePath string) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *fsHandler) handleRequest(ctx *RequestCtx) { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert

//nolint:forcetypeassert
//nolint:forcetypeassert

type byteRangeUpdater interface {
	UpdateByteRange(startPos, endPos int) error
}

func ParseByteRange(byteRange []byte, contentLength int) (startPos, endPos int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (h *fsHandler) openIndexFile(ctx *RequestCtx, dirPath string, mustCompress bool, fileEncoding string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	errDirIndexRequired   = errors.New("directory index required")
	errNoCreatePermission = errors.New("no 'create file' permissions")
)

func (h *fsHandler) createDirIndex(ctx *RequestCtx, dirPath string, mustCompress bool, fileEncoding string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	fsMinCompressRatio        = 0.8
	fsMaxCompressibleFileSize = 8 * 1024 * 1024
)

func (h *fsHandler) compressAndOpenFSFile(filePath, fileEncoding string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) compressFileNolock(
	f fs.File, fileInfo fs.FileInfo, filePath, compressedFilePath, fileEncoding string,
) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) newCompressedFSFileCache(f fs.File, fileInfo fs.FileInfo, filePath, fileEncoding string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) newCompressedFSFile(filePath, fileEncoding string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) openFSFile(filePath string, mustCompress bool, fileEncoding string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) newFSFile(f fs.File, fileInfo fs.FileInfo, compressed bool, filePath, fileEncoding string) (*fsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readFileHeader(f io.Reader, compressed bool, fileEncoding string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stripLeadingSlashes(path []byte, stripSlashes int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func hasDotDotPathSegment(path []byte) bool { _ = "STUB: not implemented"; return false }

func fileExtension(path string, compressed bool, compressedFileSuffix string) string {
	_ = "STUB: not implemented"
	return ""
}

func FileLastModified(path string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func fsModTime(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

var (
	filesLockMu  sync.Mutex
	filesLockMap = make(map[string]*fileLock)
)

type fileLock struct {
	mu sync.Mutex

	refs int
}

func acquireFileLock(absPath string) *fileLock { _ = "STUB: not implemented"; return nil }

func releaseFileLock(absPath string, flock *fileLock) { _ = "STUB: not implemented"; return }

var _ fs.FS = (*osFS)(nil)

type osFS struct{}

func (o *osFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}
func (o *osFS) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}
