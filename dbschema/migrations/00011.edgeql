CREATE MIGRATION m1cglbcapryzsb543ebdoq7yhj26qgd6p2cup5b36xc6o425ud3jka
    ONTO m1xu32j5cdchs3sahuqomufq5oyg44dtkbbxalwymurkv35i5zb6ba
{
  CREATE TYPE default::CVSection EXTENDING default::Timestamps, default::Sort {
      CREATE REQUIRED PROPERTY icon -> std::str;
      CREATE REQUIRED PROPERTY name -> std::str {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  ALTER TYPE default::Links RENAME TO default::SiteLink;
  ALTER TYPE default::Technologies RENAME TO default::CVTechnology;
};
