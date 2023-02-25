package com.kongrentian.plugins.nexus.main;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.fasterxml.jackson.dataformat.yaml.YAMLGenerator;
import com.fasterxml.jackson.datatype.joda.JodaModule;
import com.kongrentian.plugins.nexus.api.ApiClient;
import com.kongrentian.plugins.nexus.api.BundleConfigurationApi;
import com.kongrentian.plugins.nexus.api.MonitoringApi;
import com.kongrentian.plugins.nexus.api.RemoteScanApi;
import com.kongrentian.plugins.nexus.capability.SecurityCapability;
import com.kongrentian.plugins.nexus.capability.SecurityCapabilityConfiguration;
import com.kongrentian.plugins.nexus.capability.SecurityCapabilityDescriptor;
import com.kongrentian.plugins.nexus.capability.SecurityCapabilityFormField;
import com.kongrentian.plugins.nexus.model.bundle.configuration.BundleConfiguration;
import org.joda.time.DateTime;
import org.joda.time.DateTimeZone;
import org.joda.time.format.DateTimeFormat;
import org.joda.time.format.DateTimeFormatter;
import org.sonatype.nexus.capability.CapabilityReference;
import org.sonatype.nexus.capability.CapabilityRegistry;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;
import javax.inject.Inject;
import javax.inject.Named;
import javax.inject.Singleton;
import java.text.SimpleDateFormat;
import java.time.Instant;
import java.util.HashMap;
import java.util.Map;


@Named
@Singleton
public class BundleHelper {

    public final static String DATE_FORMAT_PATTERN = "yyyy-MM-dd";
    public final static SimpleDateFormat DATE_FORMAT =
            new SimpleDateFormat(DATE_FORMAT_PATTERN);
    public final static ObjectMapper MAPPER_YAML = configureObjectMapper(new ObjectMapper(
            new YAMLFactory().disable(YAMLGenerator.Feature.WRITE_DOC_START_MARKER)));


    public final static ObjectMapper MAPPER_JSON =
            configureObjectMapper(new ObjectMapper());
    private static final DateTimeFormatter dateTimeFormatter = DateTimeFormat
            .forPattern(DATE_FORMAT_PATTERN).withZone(DateTimeZone.UTC);
    private final CapabilityRegistry capabilityRegistry;
    private final Map<String, Object> capabilityStatus = new HashMap<>();
    private BundleConfiguration bundleConfiguration = new BundleConfiguration();
    private CapabilityReference securityCapabilityReference;
    private RemoteScanApi remoteScanApi;
    private MonitoringApi monitoringApi;
    private BundleConfigurationApi bundleConfigurationApi;

    @Inject
    public BundleHelper(final CapabilityRegistry capabilityRegistry) {
        this.capabilityRegistry = capabilityRegistry;
    }

    private static boolean isTypeEqual(CapabilityReference reference) {
        return SecurityCapabilityDescriptor.CAPABILITY_TYPE
                .equals(reference.context().type());
    }

    public static DateTime parseTime(String text) {
        return DateTime.parse(text, dateTimeFormatter);
    }

    public static String todayDate() {
        return DateTime.now().toString(dateTimeFormatter);
    }

    private static ObjectMapper configureObjectMapper(ObjectMapper objectMapper) {
        return objectMapper
                // if a collection (list, for example) is not present
                // during deserialization, it will be null
                // very inconvenient, causes hard to debug null exceptions
                // this makes such collections empty, instead of null
                .setDefaultSetterInfo(JsonSetter.Value.forContentNulls(Nulls.AS_EMPTY))
                // write dates as ISO dates, not as timestamps
                .disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS)
                // if a property is null, do not include it in json
                // during serialization
                .setSerializationInclusion(JsonInclude.Include.NON_NULL)
                // joda datetime handling
                .registerModule(new JodaModule())
                .setDateFormat(DATE_FORMAT);
    }

    @Nullable
    private CapabilityReference findCapability() {
        return capabilityRegistry
                .get(BundleHelper::isTypeEqual)
                .stream()
                .findFirst()
                .orElse(null);
    }

    public CapabilityReference createDefaultCapability() {
        return capabilityRegistry.add(
                SecurityCapabilityDescriptor.CAPABILITY_TYPE,
                true,
                "Automatically created at " + Instant.now().toString(),
                SecurityCapabilityFormField.createDefaultProperties());
    }

    @Nonnull
    public SecurityCapabilityConfiguration getCapabilityConfiguration() {
        return getOrCreateCapability()
                .capabilityAs(SecurityCapability.class)
                .getConfig();
    }

    @Nonnull
    public CapabilityReference getOrCreateCapability() {
        if (securityCapabilityReference != null) {
            return securityCapabilityReference;
        }
        CapabilityReference capabilityReference = findCapability();
        if (capabilityReference == null) {
            securityCapabilityReference = createDefaultCapability();
        } else {
            securityCapabilityReference = capabilityReference;
        }
        return securityCapabilityReference;
    }

    public void recreateRemoteScanApi() {
        remoteScanApi = ApiClient.createRemoteScanApi(
                bundleConfiguration,
                getCapabilityConfiguration());
    }


    public void recreateMonitoringApi() {
        monitoringApi = ApiClient.createMonitoringApi(
                bundleConfiguration,
                getCapabilityConfiguration());
    }

    public void recreateBundleConfigurationApi() {
        bundleConfigurationApi = ApiClient.createBundleConfigurationApi(
                getCapabilityConfiguration());
    }

    public boolean isCapabilityActive() {
        return getOrCreateCapability()
                .context()
                .isActive();
    }

    public RemoteScanApi getSecurityClientApi() {
        if (remoteScanApi == null) {
            recreateRemoteScanApi();
        }
        return remoteScanApi;
    }

    public MonitoringApi getMonitoringApi() {
        if (monitoringApi == null) {
            recreateMonitoringApi();
        }
        return monitoringApi;
    }

    public BundleConfigurationApi getBundleConfigurationApi() {
        if (bundleConfigurationApi == null) {
            recreateBundleConfigurationApi();
        }
        return bundleConfigurationApi;
    }

    public BundleConfiguration getBundleConfiguration() {
        return bundleConfiguration;
    }

    public void setBundleConfiguration(BundleConfiguration bundleConfiguration) {
        this.bundleConfiguration = bundleConfiguration;
        recreateRemoteScanApi();
        recreateMonitoringApi();
    }

    public Map<String, Object> getCapabilityStatus() {
        return capabilityStatus;
    }

}
