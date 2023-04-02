import React, { Key, ReactNode } from 'react';
import Typography, { TypographyProps } from '@mui/material/Typography';
import Box from '@mui/material//Box';
import { ApolloError, ServerError } from '@apollo/client';
import { useTranslation } from 'react-i18next';
import type { SxProps } from '@mui/material';

export interface Props {
  error?: ApolloError;
  errors?: ApolloError[];
  noMargin?: boolean;
  containerBoxSx?: SxProps;
  errorTitleTypographyProps?: TypographyProps;
  errorElementTypographyProps?: TypographyProps;
  ulSx?: SxProps;
  liSx?: SxProps;
}

export interface CustomNetworkError {
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

  const elements: ReactNode[] = [];

  // Loop over errors
  errorList?.forEach((err, mainIndex) => {
    const contents: { content: string; key: Key }[] = [];
    if (
      err.networkError &&
      (!(err.networkError as ServerError).result ||
        Object.keys((err.networkError as ServerError).result).length === 0 ||
        !err.graphQLErrors)
    ) {
      contents.push({ content: err.networkError.message, key: mainIndex });
    }

    if (err.graphQLErrors) {
      contents.push(
        ...err.graphQLErrors.map(({ message, extensions }, i) => {
          let mess = message;
          // Check if there is a code in extensions
          if (extensions && extensions.code) {
            mess = t(`common.errorCode.${extensions.code}`);
          }

          return { content: mess, key: `${mainIndex}-graphQLErrors-${i}` };
        }),
      );
    }

    if (
      err.networkError &&
      (err.networkError as ServerError).result &&
      (err.networkError as ServerError).result.errors
    ) {
      contents.push(
        ...((err.networkError as ServerError).result.errors as [CustomNetworkError]).map(({ message, path }, i) => ({
          content: `${path?.join('.')} ${message}`,
          key: `${mainIndex}-networkError-${i}`,
        })),
      );
    }

    // Check if a content have been detected
    if (contents) {
      // Save element
      elements.push(
        ...contents.map(({ content, key }) => (
          <Box component="li" key={key} sx={liSx}>
            <Typography color="error" {...errorElementTypographyProps}>
              {content}
            </Typography>
          </Box>
        )),
      );
    }
  });

  return (
    <Box sx={{ color: 'error', margin: noMargin ? 0 : 8, ...containerBoxSx }}>
      <Typography color="error" {...errorTitleTypographyProps}>
        {t('common.errors')}:
      </Typography>
      <Box component="ul" sx={ulSx}>
        {elements}
      </Box>
    </Box>
  );
}
/* eslint-enable */

GraphqlErrors.defaultProps = defaultProps;

export default GraphqlErrors;
