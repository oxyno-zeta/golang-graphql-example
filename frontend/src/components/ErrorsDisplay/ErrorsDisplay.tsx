import React, { Key, ReactNode } from 'react';
import Typography, { TypographyProps } from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { ApolloError, ServerError } from '@apollo/client';
import { useTranslation } from 'react-i18next';
import type { SxProps } from '@mui/material';
import {
  GraphqlErrorsExtensionsCodeCustomComponentMapKeyPrefix,
  ServerErrorCustomComponentMapKey,
  NetworkErrorCustomComponentMapKey,
} from './constants';

export interface Props {
  error?: ApolloError | Error | null;
  errors?: (ApolloError | Error)[];
  noMargin?: boolean;
  containerBoxSx?: SxProps;
  errorTitleTypographyProps?: TypographyProps;
  errorElementTypographyProps?: TypographyProps;
  ulSx?: SxProps;
  liSx?: SxProps;
  customErrorComponents?: Record<string, React.ElementType>;
  customErrorComponentProps?: Record<string, object>;
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
  customErrorComponents: {},
  customErrorComponentProps: {},
};

/* eslint-disable react/no-array-index-key */
function ErrorsDisplay({
  error,
  errors,
  noMargin,
  containerBoxSx,
  errorTitleTypographyProps,
  errorElementTypographyProps,
  ulSx,
  liSx,
  customErrorComponents = {},
  customErrorComponentProps = {},
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
    const contents: { content: string; key: Key; CustomComp?: React.ElementType; customCompProps?: object }[] = [];
    if ((err as ApolloError).networkError || (err as ApolloError).graphQLErrors) {
      // Create intermediate variable just for typescript.............
      const workingErr: ApolloError = err as ApolloError;

      if (
        workingErr.networkError &&
        (!(workingErr.networkError as ServerError).result ||
          Object.keys((workingErr.networkError as ServerError).result).length === 0 ||
          !workingErr.graphQLErrors)
      ) {
        contents.push({
          content: workingErr.networkError.message,
          key: mainIndex,
          CustomComp: customErrorComponents[ServerErrorCustomComponentMapKey],
          customCompProps: customErrorComponentProps[ServerErrorCustomComponentMapKey],
        });
      }

      if (workingErr.graphQLErrors) {
        contents.push(
          ...workingErr.graphQLErrors.map(({ message, extensions }, i) => {
            let mess = message;
            let customComp: React.ElementType | undefined;
            let customCompProps: object | undefined;
            // Check if there is a code in extensions
            if (extensions && extensions.code) {
              mess = t(`common.errorCode.${extensions.code}`);

              const mapKey = `${GraphqlErrorsExtensionsCodeCustomComponentMapKeyPrefix}${extensions.code}`;

              customComp = customErrorComponents[mapKey];
              customCompProps = customErrorComponentProps[mapKey];
            }

            return {
              content: mess,
              key: `${mainIndex}-graphQLErrors-${i}`,
              CustomComp: customComp,
              customCompProps,
            };
          }),
        );
      }

      if (
        workingErr.networkError &&
        (workingErr.networkError as ServerError).result &&
        ((workingErr.networkError as ServerError).result as Record<string, [CustomNetworkError]>).errors &&
        Array.isArray(((workingErr.networkError as ServerError).result as Record<string, [CustomNetworkError]>).errors)
      ) {
        contents.push(
          ...((workingErr.networkError as ServerError).result as Record<string, [CustomNetworkError]>).errors
            .filter((it) => typeof it === 'object' && it !== null) // Ensure that it isn't null or not an object
            .map(({ message, path }, i) => ({
              content: `${path?.join('.')} ${message}`,
              key: `${mainIndex}-networkError-${i}`,
              CustomComp: customErrorComponents[NetworkErrorCustomComponentMapKey],
              customCompProps: customErrorComponentProps[NetworkErrorCustomComponentMapKey],
            })),
        );
      }
    } else {
      // Isn't an apollo error
      contents.push({ content: err.message, key: mainIndex });
    }

    if (contents) {
      // Check if a content have been detected
      // Save element
      elements.push(
        ...contents.map(({ content, key, CustomComp, customCompProps }) => (
          <Box component="li" key={key} sx={liSx}>
            {CustomComp ? (
              <CustomComp {...customCompProps} />
            ) : (
              <Typography color="error" {...errorElementTypographyProps}>
                {content}
              </Typography>
            )}
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

ErrorsDisplay.defaultProps = defaultProps;

export default ErrorsDisplay;
