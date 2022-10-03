import React, { Fragment } from 'react';
import Typography from '@mui/material/Typography';
import Box from '@mui/material//Box';
import { ApolloError, ServerError } from '@apollo/client';
import { useTranslation } from 'react-i18next';

interface Props {
  error?: ApolloError;
  errors?: ApolloError[];
  noMargin?: boolean;
}

interface CustomNetworkError {
  message: string;
  path: null | undefined | ReadonlyArray<string | number>;
}

const defaultProps = {
  error: null,
  errors: [],
  noMargin: false,
};

/* eslint-disable react/no-array-index-key */
function GraphqlErrors({ error, errors, noMargin }: Props) {
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
    <Box sx={{ color: 'error', margin: noMargin ? 0 : 8 }}>
      <Typography color="error">{t('common.errors')}:</Typography>
      <ul>
        {errorList?.map((err, mainIndex) => (
          <Fragment key={mainIndex}>
            {err.networkError && (!(err.networkError as ServerError).result || !err.graphQLErrors) && (
              <li key={mainIndex}>
                <Typography color="error">{err.networkError.message}</Typography>
              </li>
            )}
            {err.graphQLErrors &&
              err.graphQLErrors.map(({ message, extensions }, i) => {
                let mess = message;
                // Check if there is a code in extensions
                if (extensions.code) {
                  mess = t(`common.errorCode.${extensions.code}`);
                }

                return (
                  <li key={`${mainIndex}-graphQLErrors-${i}`}>
                    <Typography color="error">{mess}</Typography>
                  </li>
                );
              })}
            {err.networkError &&
              (err.networkError as ServerError).result &&
              (err.networkError as ServerError).result.errors &&
              ((err.networkError as ServerError).result.errors as [CustomNetworkError]).map(({ message, path }, i) => (
                <li key={`${mainIndex}-networkError-${i}`}>
                  <Typography color="error">
                    {path?.join('.')} {message}
                  </Typography>
                </li>
              ))}
          </Fragment>
        ))}
      </ul>
    </Box>
  );
}
/* eslint-enable */

GraphqlErrors.defaultProps = defaultProps;

export default GraphqlErrors;
