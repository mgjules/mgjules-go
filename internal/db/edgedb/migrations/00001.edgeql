CREATE MIGRATION m1sqtxlp32sbgozlixfea5d42c2aazje2chl6rcheswn7gi7iqrxsa
    ONTO initial
{
  CREATE ABSTRACT TYPE default::Timestamps {
      CREATE REQUIRED PROPERTY created_at -> std::datetime {
          SET default := (std::datetime_current());
      };
  };
  CREATE SCALAR TYPE default::Gender EXTENDING enum<Male, Female, Robot>;
  CREATE SCALAR TYPE default::Lang EXTENDING enum<En, Fr>;
  CREATE TYPE default::Meta EXTENDING default::Timestamps {
      CREATE REQUIRED PROPERTY base_url -> std::str;
      CREATE REQUIRED PROPERTY description -> std::str;
      CREATE REQUIRED PROPERTY first_name -> std::str;
      CREATE REQUIRED PROPERTY gender -> default::Gender;
      CREATE REQUIRED PROPERTY github -> std::str;
      CREATE REQUIRED PROPERTY keywords -> array<std::str>;
      CREATE REQUIRED PROPERTY lang -> default::Lang;
      CREATE REQUIRED PROPERTY last_name -> std::str;
      CREATE REQUIRED PROPERTY username -> std::str;
  };
};
