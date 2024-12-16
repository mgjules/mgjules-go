CREATE MIGRATION m1q2vbhjxpatvt7yc6sojvdqhojio4ojrrrjzfl4l5ogbfj2gio7oa
    ONTO m1az7gvdblyl665htlm6bknokqhvjhgerwtpv7dpj2q2yzw2my25oq
{
  CREATE TYPE default::CVLanguage EXTENDING default::Timestamps, default::Sort {
      CREATE REQUIRED PROPERTY icon -> std::str;
      CREATE REQUIRED PROPERTY level -> std::str;
      CREATE REQUIRED PROPERTY name -> std::str;
  };
};
