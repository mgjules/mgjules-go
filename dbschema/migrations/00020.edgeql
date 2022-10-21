CREATE MIGRATION m1247jhpvn2robqnl57e7n6ehdjc2hg4oeubvkt4vsaunjshie4wpq
    ONTO m1jm33r7e7rdbuog2nlkecqcspu44ew4lo4uaye7jpsov4ntchlr5q
{
  CREATE ABSTRACT TYPE default::Period {
      CREATE REQUIRED PROPERTY from -> cal::local_date;
      CREATE PROPERTY to -> cal::local_date;
  };
  CREATE SCALAR TYPE default::Role EXTENDING enum<Founder, `Co-Founder`, Speaker, `Co-Speaker`>;
  CREATE TYPE default::CVContribution EXTENDING default::Timestamps, default::Period {
      CREATE REQUIRED PROPERTY event -> std::str;
      CREATE REQUIRED PROPERTY link -> std::str;
      CREATE REQUIRED PROPERTY role -> default::Role;
      CREATE REQUIRED PROPERTY title -> std::str;
  };
  ALTER TYPE default::CVExperience EXTENDING default::Period LAST;
  ALTER TYPE default::CVExperience {
      ALTER PROPERTY from {
          RESET OPTIONALITY;
          DROP OWNED;
          RESET TYPE;
      };
  };
  ALTER TYPE default::CVExperience {
      ALTER PROPERTY to {
          DROP OWNED;
          RESET TYPE;
      };
  };
};
