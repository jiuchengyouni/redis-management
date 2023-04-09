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
	export class Email {
	    address: string;
	
	    static createFrom(source: any = {}) {
	        return new Email(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = source["address"];
	    }
	}
	export class HashAddOrUpdateFieldRequest {
	    connection_id: string;
	    db: number;
	    key: string;
	    field: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new HashAddOrUpdateFieldRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.field = source["field"];
	        this.value = source["value"];
	    }
	}
	export class HashFieldDeleteRequest {
	    connection_id: string;
	    db: number;
	    key: string;
	    field: string[];
	
	    static createFrom(source: any = {}) {
	        return new HashFieldDeleteRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.field = source["field"];
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
	export class ListValueRequest {
	    connection_id: string;
	    db: number;
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new ListValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class SetValueRequest {
	    connection_id: string;
	    db: number;
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new SetValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.value = source["value"];
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
	export class ZSetValueRequest {
	    connection_id: string;
	    db: number;
	    key: string;
	    score: number;
	    member: any;
	
	    static createFrom(source: any = {}) {
	        return new ZSetValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.db = source["db"];
	        this.key = source["key"];
	        this.score = source["score"];
	        this.member = source["member"];
	    }
	}

}

