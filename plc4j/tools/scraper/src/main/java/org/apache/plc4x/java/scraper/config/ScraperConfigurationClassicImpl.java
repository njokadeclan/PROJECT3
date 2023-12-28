/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.scraper.config;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import org.apache.plc4x.java.scraper.ScrapeJob;
import org.apache.plc4x.java.scraper.config.triggeredscraper.ScraperConfigurationTriggeredImpl;
import org.apache.plc4x.java.scraper.exception.ScraperConfigurationException;
import org.apache.plc4x.java.scraper.exception.ScraperException;
import org.apache.plc4x.java.scraper.ScraperImpl;

import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

/**
 * Configuration class for {@link ScraperImpl}.
 *
 * @deprecated Scraper is deprecated please use {@link ScraperConfigurationTriggeredImpl} instead all functions are supplied as well see java-doc of {@link org.apache.plc4x.java.scraper.triggeredscraper.TriggeredScraperImpl}
 */
@Deprecated
public class ScraperConfigurationClassicImpl implements ScraperConfiguration {

    private final Map<String, String> sources;
    private final List<JobConfigurationImpl> jobConfigurations;

    /**
     * Default constructor.
     *
     * @param sources           Map from connection alias to connection string
     * @param jobConfigurations List of configurations one for each Job
     */
    @JsonCreator
    public ScraperConfigurationClassicImpl(@JsonProperty(value = "sources", required = true) Map<String, String> sources,
                                           @JsonProperty(value = "jobs", required = true) List<JobConfigurationImpl> jobConfigurations) {
        checkNoUnreferencedSources(sources, jobConfigurations);
        this.sources = sources;
        this.jobConfigurations = jobConfigurations;
    }

    private void checkNoUnreferencedSources(Map<String, String> sources, List<JobConfigurationImpl> jobConfigurations) {
        Set<String> unreferencedSources = jobConfigurations.stream()
            .flatMap(job -> job.getSources().stream())
            .filter(source -> !sources.containsKey(source))
            .collect(Collectors.toSet());
        if (!unreferencedSources.isEmpty()) {
            throw new ScraperConfigurationException("There are the following unreferenced sources: " + unreferencedSources);
        }
    }

    @Override
    public Map<String, String> getSources() {
        return sources;
    }

    @Override
    public List<JobConfigurationImpl> getJobConfigurations() {
        return jobConfigurations;
    }

    @Override
    public List<ScrapeJob> getJobs() throws ScraperException {
        return ScraperConfigurationTriggeredImpl.getJobs(jobConfigurations,sources);
    }


}
