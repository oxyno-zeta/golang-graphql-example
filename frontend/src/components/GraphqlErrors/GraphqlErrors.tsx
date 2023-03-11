import React, { Fragment } from 'react';
import Typography, { TypographyProps } from '@mui/material/Typography';
import Box from '@mui/material//Box';
import { ApolloError, ServerError } from '@apollo/client';
import { useTranslation } from 'react-i18next';
import type { SxProps } from '@mui/material';

interface Props {
  error?: ApolloError;
  errors?: ApolloError[];
  noMargin?: boolean;
  containerBoxSx?: SxProps;
  errorTitleTypographyProps?: TypographyProps;
  errorElementTypographyProps?: TypographyProps;
  ulSx?: SxProps;
  liSx?: SxProps;
}

interface CustomNetworkError {
  message: string;
  path: null | undefined | ReadonlyArray<string | number>;
}

const defaultProps = {
  error: null,
  errors: [],
  noMargin: false,
  containerBoxSx: {},
  errorTitleTypographyProps: {},
  errorElementTypographyProps: {},
  ulSx: {},
  liSx: {},
};

/* eslint-disable react/no-array-index-key */
function GraphqlErrors({
  error,
  errors,
  noMargin,
  containerBoxSx,
  errorTitleTypographyProps,
  errorElementTypographyProps,
  ulSx,
  liSx,
}: Props) {
  // Initialize translate
  const { t } = useTranslation();

  // Check if error or errors are set
  if ((!error || !errors) && errors && errors.length === 0) {
    // Display nothing
    return null;
  }

  // Build error list
  let errorList = errors;
  // Check if error is set
  if (error) {
    // Override array
    errorList = [error];
  }

  return (
    <Box sx={{ color: 'error', margin: noMargin ? 0 : 8, ...containerBoxSx }}>
      <Typography color="error" {...errorTitleTypographyProps}>
        {t('common.errors')}:
      </Typography>
      <Box component="ul" sx={ulSx}>
        {errorList?.map((err, mainIndex) => (
          <Fragment key={mainIndex}>
            {err.networkError && (!(err.networkError as ServerError).result || !err.graphQLErrors) && (
              <Box component="li" key={mainIndex} sx={liSx}>
                <Typography color="error" {...errorElementTypographyProps}>
                  {err.networkError.message}
                </Typography>
              </Box>
            )}
            {err.graphQLErrors &&
              err.graphQLErrors.map(({ message, extensions }, i) => {
                let mess = message;
                // Check if there is a code in extensions
                if (extensions && extensions.code) {
                  mess = t(`common.errorCode.${extensions.code}`);
                }

                return (
                  <Box component="li" key={`${mainIndex}-graphQLErrors-${i}`} sx={liSx}>
                    <Typography color="error" {...errorElementTypographyProps}>
                      {mess}
                    </Typography>
                  </Box>
                );
              })}
            {err.networkError &&
              (err.networkError as ServerError).result &&
              (err.networkError as ServerError).result.errors &&
              ((err.networkError as ServerError).result.errors as [CustomNetworkError]).map(({ message, path }, i) => (
                <Box component="li" key={`${mainIndex}-networkError-${i}`} sx={liSx}>
                  <Typography color="error" {...errorElementTypographyProps}>
                    {path?.join('.')} {message}
                  </Typography>
                </Box>
              ))}
          </Fragment>
        ))}
      </Box>
    </Box>
  );
}
/* eslint-enable */

GraphqlErrors.defaultProps = defaultProps;

export default GraphqlErrors;
