CREATE MIGRATION m1az7gvdblyl665htlm6bknokqhvjhgerwtpv7dpj2q2yzw2my25oq
    ONTO m1qvasiycfj6pozfrzvz3hazx3xabv5iw5qqsgt5qiijfnvx4evcsq
{
  CREATE TYPE default::CVInterest EXTENDING default::Timestamps, default::Sort {
      CREATE REQUIRED PROPERTY image -> std::str;
      CREATE REQUIRED PROPERTY name -> std::str;
  };
};
