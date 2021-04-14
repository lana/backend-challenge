package xyz.lana.challenge.robertom.api;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.MissingServletRequestParameterException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.client.RestClientException;
import xyz.lana.challenge.robertom.api.resource.ErrorResource;
import xyz.lana.challenge.robertom.exceptions.NotFoundException;

import java.security.InvalidParameterException;

@RestControllerAdvice
@Slf4j
public class ControllerAdvice {

    @ResponseBody
    @ExceptionHandler(RestClientException.class)
    @ResponseStatus(HttpStatus.SERVICE_UNAVAILABLE)
    public ErrorResource serviceUnavailableExceptionHandler(RestClientException e) {
        return generateExceptionMessage(e);
    }

    @ResponseBody
    @ExceptionHandler(NotFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    public ErrorResource handleNotFoundException(NotFoundException e) {
        return generateExceptionMessage(e);
    }

    @ResponseBody
    @ExceptionHandler(InvalidParameterException.class)
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    public ErrorResource handleInvalidParameterException(InvalidParameterException e) {
        return generateExceptionMessage(e);
    }

    @ResponseBody
    @ExceptionHandler(MethodArgumentNotValidException.class)
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    public ErrorResource handleInvalidParameterException(MethodArgumentNotValidException e) {
        return generateExceptionMessage(e);
    }

    @ResponseBody
    @ExceptionHandler(MissingServletRequestParameterException.class)
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    public ErrorResource handleMissingServletRequestParameterException(MissingServletRequestParameterException e) {
        return generateExceptionMessage(e);
    }

    /**
     * Response internal service exception
     */
    @ResponseBody
    @ExceptionHandler(Exception.class)
    @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
    public ErrorResource handleGeneralException(Exception e) {
        return generateExceptionMessage(e);
    }

    private ErrorResource generateExceptionMessage(final Exception e) {
        return generateExceptionMessage(e, false);
    }

    private ErrorResource generateExceptionMessage(final Exception e, final boolean simpleLog) {
        String logMessage = e.getMessage();

        if (logMessage == null) {
            logMessage = e.getClass().getName();
        }

        if (simpleLog) {
            log.error("############  Message : {} ############", logMessage);
        } else {
            log.error("############  Message : {} ############", logMessage, e);
        }

        ErrorResource response = new ErrorResource();
        response.setMessage(logMessage);

        return response;
    }

}