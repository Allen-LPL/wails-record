export namespace main {
	
	export class Data {
	    date: string;
	    contesta: string;
	    contestb: string;
	    typea: string;
	    typeb: string;
	    wina: string;
	    winb: string;
	    drawa: string;
	    drawb: string;
	    losea: string;
	    loseb: string;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date = source["date"];
	        this.contesta = source["contesta"];
	        this.contestb = source["contestb"];
	        this.typea = source["typea"];
	        this.typeb = source["typeb"];
	        this.wina = source["wina"];
	        this.winb = source["winb"];
	        this.drawa = source["drawa"];
	        this.drawb = source["drawb"];
	        this.losea = source["losea"];
	        this.loseb = source["loseb"];
	    }
	}

}

