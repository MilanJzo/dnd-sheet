import { createClient } from "@libsql/client";
import * as dotenv from "dotenv";
import { drizzle } from "drizzle-orm/libsql";
import feats from "./scraper/feats/formated_feats.json";
// https://www.dandwiki.com/wiki/3.5e_Homebrew_Feats

type Feat = {
    name: string;
    description: string;
    prerequisites: string;
    benefit: string;
    normal: string;
    special: string;
    link: string;
};

dotenv.config();

export const client = createClient({
    url: process.env.DATABASE_URL ?? "",
    authToken: process.env.DATABASE_AUTH_TOKEN
});
// export const client = createClient({ url: "file:./drizzle/db.sqlite" });
export const db = drizzle(client);

feats.forEach(async (feat) => {
    // await db.feats.insert({
    //     name: feat.name,
    //     description: feat.description,
    //     prerequisites: feat.prerequisites,
    //     benefit: feat.benefit,
    //     normal: feat.normal,
    //     special: feat.special,
    //     link: feat.link
    // });
});
