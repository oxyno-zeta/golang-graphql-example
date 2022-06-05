import React from 'react';
import Typography from '@mui/material/Typography';
import Box from '@mui/material//Box';
import { ApolloError, ServerError } from '@apollo/client';
import { useTranslation } from 'react-i18next';

interface Props {
  error: ApolloError;
  noMargin?: boolean;
}

interface CustomNetworkError {
  message: string;
  path: null | undefined | ReadonlyArray<string | number>;
}

/* eslint-disable react/no-array-index-key */
function GraphqlErrors({ error, noMargin }: Props) {
  // Initialize translate
  const { t } = useTranslation();

  return (
    <Box sx={{ color: 'error', margin: noMargin ? 0 : 8 }}>
      <Typography color="error">{t('common.errors')}:</Typography>
      <ul>
        {error.networkError && (!(error.networkError as ServerError).result || !error.graphQLErrors) && (
          <li>
            <Typography color="error">{error.networkError.message}</Typography>
          </li>
        )}
        {error.graphQLErrors &&
          error.graphQLErrors.map(({ message }, i) => (
            <li key={`graphQLErrors-${i}`}>
              <Typography color="error">{message}</Typography>
            </li>
          ))}
        {error.networkError &&
          (error.networkError as ServerError).result &&
          (error.networkError as ServerError).result.errors &&
          ((error.networkError as ServerError).result.errors as [CustomNetworkError]).map(({ message, path }, i) => (
            <li key={`networkError-${i}`}>
              <Typography color="error">
                {path?.join('.')} {message}
              </Typography>
            </li>
          ))}
      </ul>
    </Box>
  );
}
/* eslint-enable */

GraphqlErrors.defaultProps = {
  noMargin: false,
};

export default GraphqlErrors;
