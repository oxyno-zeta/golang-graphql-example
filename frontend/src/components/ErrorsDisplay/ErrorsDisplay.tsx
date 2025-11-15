import React, { type Key, type ReactNode } from 'react';
import Typography, { type TypographyProps } from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { ServerError, CombinedGraphQLErrors } from '@apollo/client';
import { useTranslation } from 'react-i18next';
import type { SxProps } from '@mui/material';
import { isAxiosError, type AxiosError } from 'axios';
import { type GraphQLFormattedErrorExtensions } from 'graphql';
import WithTraceError from '~utils/WithTraceError';
import { GraphqlErrorsExtensionsCodeCustomComponentMapKeyPrefix, ServerErrorCustomComponentMapKey } from './constants';

export interface Props {
  readonly error?: Error | null;
  readonly errors?: Error[];
  readonly noMargin?: boolean;
  readonly containerBoxSx?: SxProps;
  readonly errorTitleTypographyProps?: TypographyProps;
  readonly errorElementTypographyProps?: TypographyProps;
  readonly ulSx?: SxProps;
  readonly liSx?: SxProps;
  readonly customErrorComponents?: Record<string, React.ElementType>;
  readonly customErrorComponentProps?: Record<string, object>;
}

export interface CustomNetworkError {
  message: string;
  path: null | undefined | readonly (string | number)[];
}

interface ErrorRESTBody {
  error: string;
  extensions: GraphQLFormattedErrorExtensions | undefined;
}
interface ErrorContent {
  content: string;
  key: Key;
  CustomComp?: React.ElementType;
  customCompProps?: object;
}

function errorBodyToErrorContent(
  t: (input: string) => string,
  message: string,
  extensions: GraphQLFormattedErrorExtensions | undefined,
  key: string,
  customErrorComponents: Record<string, React.ElementType>,
  customErrorComponentProps: Record<string, object>,
) {
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
    key,
    CustomComp: customComp,
    customCompProps,
  };
}

function ErrorsDisplay({
  error = null,
  errors = [],
  noMargin = false,
  containerBoxSx = {},
  errorTitleTypographyProps = {},
  errorElementTypographyProps = {},
  ulSx = {},
  liSx = {},
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
  errorList?.forEach((errOr, mainIndex) => {
    let err = errOr;
    let requestId = '';
    let traceId = '';
    // Check if it is a WithTraceError
    if (WithTraceError.is(err)) {
      // Replace error and save trace and request id information
      traceId = err.traceId;
      requestId = err.requestId;
      // Erase error
      err = err.err;
    }

    const contents: ErrorContent[] = [];

    // Check if it is an AxiosError
    if (isAxiosError(err)) {
      const axErr = err as AxiosError<ErrorRESTBody>;
      if (axErr.response && axErr.response.data && axErr.response.data.error) {
        const body = axErr.response.data;

        contents.push(
          errorBodyToErrorContent(
            t,
            body.error,
            body.extensions,
            `${mainIndex}-axiosError`,
            customErrorComponents,
            customErrorComponentProps,
          ),
        );
      } else {
        // Default
        contents.push({ content: err.message, key: mainIndex });
      }
    } else if (ServerError.is(err) || CombinedGraphQLErrors.is(err)) {
      if (ServerError.is(err)) {
        contents.push({
          content: err.message,
          key: mainIndex,
          CustomComp: customErrorComponents[ServerErrorCustomComponentMapKey],
          customCompProps: customErrorComponentProps[ServerErrorCustomComponentMapKey],
        });
      }

      if (CombinedGraphQLErrors.is(err)) {
        contents.push(
          ...err.errors.map(({ message, extensions }, i) =>
            errorBodyToErrorContent(
              t,
              message,
              extensions,
              `${mainIndex}-graphQLErrors-${i}`,
              customErrorComponents,
              customErrorComponentProps,
            ),
          ),
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
              <CustomComp content={content} error={err} traceId={traceId} requestId={requestId} {...customCompProps} />
            ) : (
              <>
                <Typography color="error" {...errorElementTypographyProps}>
                  {content}
                </Typography>
                {traceId ? (
                  <Typography component="p" color="error" variant="caption" {...errorElementTypographyProps}>
                    {t('common.supportTraceId')}: {traceId}
                  </Typography>
                ) : null}
                {requestId ? (
                  <Typography component="p" color="error" variant="caption" {...errorElementTypographyProps}>
                    {t('common.supportRequestId')}: {requestId}
                  </Typography>
                ) : null}
              </>
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

export default ErrorsDisplay;
