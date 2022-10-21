CREATE MIGRATION m1c7q5w5cd6u57gcl6kirqji3zxfty2ypwbsxjcswjc5bgvc6hdnxq
    ONTO m1sqtxlp32sbgozlixfea5d42c2aazje2chl6rcheswn7gi7iqrxsa
{
  ALTER TYPE default::Meta {
      CREATE REQUIRED PROPERTY avatar -> std::str {
          SET default := 'https://avatars.githubusercontent.com/u/38979769?v=4';
      };
  };
};
