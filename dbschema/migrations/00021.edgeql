CREATE MIGRATION m1qvasiycfj6pozfrzvz3hazx3xabv5iw5qqsgt5qiijfnvx4evcsq
    ONTO m1247jhpvn2robqnl57e7n6ehdjc2hg4oeubvkt4vsaunjshie4wpq
{
  CREATE TYPE default::CVAward EXTENDING default::Timestamps {
      CREATE REQUIRED PROPERTY date -> cal::local_date;
      CREATE REQUIRED PROPERTY description -> std::str;
      CREATE REQUIRED PROPERTY event -> std::str;
      CREATE REQUIRED PROPERTY icon -> std::str;
      CREATE REQUIRED PROPERTY link -> std::str;
      CREATE REQUIRED PROPERTY result -> std::str;
  };
};
