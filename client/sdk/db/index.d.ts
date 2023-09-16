/**
 *
 * @param {{url?:string, apikey?:string}} param0
 */
export function connect({ url, apikey }: {
    url?: string;
    apikey?: string;
}): {
    select: typeof select;
    insert: typeof insert;
    update: typeof update;
    delete: typeof db_delete;
    addEventListener: typeof addEventListener;
};
/**
 *
 * @param {string[]|"*"} rows
 */
declare function select(rows: string[] | "*"): {
    from: (table: string) => {
        where: (condition: string | null) => {
            query: () => Promise<any>;
        };
        query: () => Promise<any>;
    };
};
/**
 *
 * @param {{table:string, values:{}}}} param0
 */
declare function insert({ table, values }: {
    table: string;
    values: {};
}): void;
/**
 *
 * @param {{query:string, values:{[any:string]:any}}}} param0
 */
declare function update({ table, query, values }: {
    query: string;
    values: {
        [any: string]: any;
    };
}): void;
/**
 * @param {Object} param0
 * @param {string} param0.table
 * @param {string} param0.query
 */
declare function db_delete({ table, query }: {
    table: string;
    query: string;
}): void;
/**
 * @param {"insert"|"update"|"delete"} event
 * @param {*} f
 */
declare function addEventListener(table: any, event: "insert" | "update" | "delete", f: any): void;
export {};
