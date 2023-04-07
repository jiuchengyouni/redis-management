export namespace define {
	
	export class Connection {
	    id: string;
	    name: string;
	    address: string;
	    port: string;
	    username: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.address = source["address"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	    }
	}
	export class CreateKeyValueRequest {
	    connection_id: string;
	    db: number;
	    key: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateKeyValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.type = source["type"];
	    }
	}
	export class KeyListRequest {
	    connection_id: string;
	    db: number;
	    keyword: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyListRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.keyword = source["keyword"];
	    }
	}
	export class KeyValueRequest {
	    connection_id: string;
	    db: number;
	    key: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.key = source["key"];
	    }
	}
	export class UpdateKeyValueRequest {
	    connection_id: string;
	    db: number;
	    key: string;
	    ttl: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateKeyValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.ttl = source["ttl"];
	        this.value = source["value"];
	    }
	}

}

