CREATE MIGRATION m1ec7wgoq6fnxmeo6qhqdrofekko2ttgh435kxvnn2f3jqcvh7bb7a
    ONTO m1cglbcapryzsb543ebdoq7yhj26qgd6p2cup5b36xc6o425ud3jka
{
  CREATE TYPE default::CVExperience EXTENDING default::Timestamps {
      CREATE MULTI LINK technologies -> default::CVTechnology {
          CREATE PROPERTY sort -> std::int32 {
              SET default := 0;
          };
      };
      CREATE REQUIRED PROPERTY company -> std::str;
      CREATE REQUIRED PROPERTY from -> cal::local_date;
      CREATE REQUIRED PROPERTY link -> std::str;
      CREATE REQUIRED PROPERTY tasks -> array<std::str>;
      CREATE PROPERTY to -> cal::local_date;
  };
};
