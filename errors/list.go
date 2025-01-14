package errors

var (
	ApplicationNotSet    = New("application instance is not initialized")
	ArtisanFacadeNotSet  = New("artisan facade is not initialized")
	CacheFacadeNotSet    = New("cache facade is not initialized")
	ConfigFacadeNotSet   = New("config facade is not initialized")
	JSONParserNotSet     = New("JSON parser is not initialized")
	LogFacadeNotSet      = New("log facade is not initialized")
	OrmFacadeNotSet      = New("orm facade is not initialized")
	QueueFacadeNotSet    = New("queue facade is not initialized")
	ScheduleFacadeNotSet = New("schedule facade is not initialized")
	StorageFacadeNotSet  = New("storage facade is not initialized")
	InvalidHttpContext   = New("invalid http context")

	ConsoleFileAlreadyExists = New("the %s already exists. Use the --force or -f flag to overwrite").SetModule(ModuleConsole)
	ConsoleEmptyFieldValue   = New("the %s name cannot be empty").SetModule(ModuleConsole)

	DockerUnknownContainerType           = New("unknown container type")
	DockerInsufficientDatabaseContainers = New("the number of database container is not enough, expect: %d, got: %d")
	DockerDatabaseContainerCountZero     = New("the number of database container must be greater than 0")
	DockerMissingContainerId             = New("no container id return when creating %s docker")

	AuthEmptySecret         = New("authentication secret is missing or required")
	AuthInvalidClaims       = New("authentication token contains invalid claims")
	AuthInvalidKey          = New("authentication key is invalid")
	AuthInvalidToken        = New("authentication token is invalid")
	AuthNoPrimaryKeyField   = New("no primary key field found in the model, ensure primary key is set, e.g., orm.Model")
	AuthParseTokenFirst     = New("authentication token must be parsed first")
	AuthRefreshTimeExceeded = New("authentication refresh time limit exceeded")
	AuthTokenDisabled       = New("authentication token has been disabled")
	AuthTokenExpired        = New("authentication token has expired")

	CacheDriverNotSupported        = New("invalid driver: %s, only support memory, custom")
	CacheForeverFailed             = New("cache forever is failed")
	CacheMemoryInvalidIntValueType = New("value type of %s is not *atomic.Int64 or *int64 or *atomic.Int32 or *int32")
	CacheStoreContractNotFulfilled = New("%s doesn't implement contracts/cache/store")
	CacheSupportRequired           = New("cache support is required")

	CryptAppKeyNotSet        = New("APP_KEY is required in artisan environment")
	CryptInvalidAppKeyLength = New("invalid APP_KEY length: %d bytes")
	CryptMissingIVKey        = New("decrypt payload error: missing iv key")
	CryptMissingValueKey     = New("decrypt payload error: missing value key")

	DBForceIsRequiredInProduction = New("application in production use --force to run this command")
	DBSeederNotFound              = New("not found %s seeder")

	EventListenerNotBind = New("event %v doesn't bind listeners")

	FilesystemDefaultDiskNotSet   = New("please set default disk")
	FilesystemDeleteDirectory     = New("can't delete directory, please use DeleteDirectory")
	FilesystemDriverNotSupported  = New("invalid driver: %s, only support local, custom")
	FilesystemFileNotExist        = New("file doesn't exist")
	FilesystemInvalidCustomDriver = New("init %s disk fail: via must be implement filesystem.Driver or func() (filesystem.Driver, error)")

	GrpcEmptyClientHost         = New("client's host can't be empty")
	GrpcEmptyClientPort         = New("client's port can't be empty")
	GrpcEmptyServerHost         = New("host can't be empty")
	GrpcEmptyServerPort         = New("port can't be empty")
	GrpcInvalidInterceptorsType = New("the type of clients.%s.interceptors must be []string")

	LangFileNotExist = New("translation file does not exist")

	LogDriverCircularReference = New("%s driver can't include self channel").SetModule(ModuleLog)
	LogDriverNotSupported      = New("invalid driver: %s, only support stack, single, daily, custom").SetModule(ModuleLog)
	LogEmptyLogFilePath        = New("empty log file path").SetModule(ModuleLog)

	MigrationNameIsRequired    = New("migration name cannot be empty")
	MigrationUnsupportedDriver = New("unsupported migration driver: %s")

	OrmDatabaseConfigNotFound      = New("not found database configuration")
	OrmDriverNotSupported          = New("invalid driver: %s, only support mysql, postgres, sqlite and sqlserver")
	OrmFailedToGenerateDNS         = New("failed to generate DSN for connection: %s")
	OrmFactoryMissingAttributes    = New("failed to get raw attributes")
	OrmFactoryMissingMethod        = New("%s does not find factory method")
	OrmInitConnection              = New("init %s connection error: %v")
	OrmMissingWhereClause          = New("WHERE conditions required")
	OrmNoDialectorsFound           = New("no dialectors found")
	OrmQueryAssociationsConflict   = New("cannot set orm.Associations and other fields at the same time")
	OrmQueryConditionRequired      = New("query condition is required")
	OrmQueryEmptyId                = New("id can't be empty")
	OrmQueryEmptyRelation          = New("relation can't be empty")
	OrmQueryInvalidModel           = New("invalid model %s")
	OrmQueryInvalidParameter       = New("parameter error, please check the document")
	OrmQueryModelNotPointer        = New("model must be pointer")
	OrmQuerySelectAndOmitsConflict = New("cannot set Select and Omits at the same time")
	OrmRecordNotFound              = New("record not found")
	OrmUnexpectedQueryType         = New("unexpected Query type %T, expected *gorm.Query")

	QueueDriverNotSupported     = New("unknown queue driver: %s")
	QueueDuplicateJobSignature  = New("job signature duplicate: %s, the names of Job and Listener cannot be duplicated")
	QueueEmptyJobSignature      = New("the Signature of job can't be empty")
	QueueEmptyListenerSignature = New("the Signature of listener can't be empty")

	RouteDefaultDriverNotSet = New("please set default driver")
	RouteInvalidDriver       = New("init %s route driver fail: route must be implement route.Route or func() (route.Route, error)")

	SchemaDriverNotSupported = New("driver %s is not supported")
	SchemaFailedToGetTables  = New("failed to get %s tables: %v")

	SessionDriverAlreadyExists   = New("session driver [%s] already exists")
	SessionDriverExtensionFailed = New("session failed to extend session [%s] driver [%v]")
	SessionDriverIsNotSet        = New("session driver is not set")
	SessionDriverNotSupported    = New("session driver [%s] not supported")
	SessionNotFound              = New("session [%s] not found")

	UnknownFileExtension = New("unknown file extension")

	ValidationDataInvalidType = New("data must be map[string]any or map[string][]string or struct")
	ValidationDuplicateFilter = New("duplicate filter name: %s")
	ValidationDuplicateRule   = New("duplicate rule name: %s")
	ValidationEmptyData       = New("data can't be empty")
	ValidationEmptyRules      = New("rules can't be empty")
)
