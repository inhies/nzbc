package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/inhies/newznab/client"
	"github.com/kylelemons/godebug/pretty"
)

var Pretty = pretty.Config{PrintStringers: true}
var producers, consumers sync.WaitGroup

const nzns = "http://www.newznab.com/DTD/2010/feeds/attributes/"

// This will eventually lex and parse a query to enable the use of search
// operators.
func parseQuery(query string) (*client.SearchRequest, error) {
	var request = new(client.SearchRequest)
	request.Query = strings.Join(os.Args[1:], " ")

	return request, nil
}
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please specify a search term")
	}

	request, err := parseQuery(strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Println(err)
		return
	}
	//request.Limit = 4

	for _, indexer := range Conf.Indexers {
		if !indexer.Enabled {
			continue
		}
		fmt.Println("searching", indexer.Name, "...")
		results, err := fetch(indexer, request)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var i = indexer
		var count = len(results.NZBs)
		fmt.Println(i.Name,
			"returned results", request.Offset+1,
			"to", request.Offset+count,
			"out of", results.Total)
		for rI, item := range results.NZBs {
			//Pretty.Print(item)
			duration := time.Since(item.PubDate.Time).Hours() / 24
			fmt.Println(rI+request.Offset+1, "/", results.Total, ":",
				item.Category, ":",
				item.Title, Round(duration, 0, 0), "days old")
			fmt.Println(item.Link, "\n")
		}

		/*
		 */
	}
}

func fetch(serv client.Indexer, request *client.SearchRequest) (*client.SearchResults, error) {

	results, err := serv.Search(request)
	if err != nil {
		return nil, fmt.Errorf(
			"Error searching for \"%s\": from \"%s\": %s", request.Query, serv.Name, err)
	}
	if results != nil {
		if results.Error != nil {
			return results, nil
		}
		return results, nil
		/*
			count := len(results.NZBs)
			if count < results.Total {
				request.Offset += count
				//go fetch(serv, request, Consolidator.add(make(chan Results)), false)
			}
		*/
	}
	return nil, fmt.Errorf("No results returned")
}

var testdata = []byte(`
<rss
    xmlns:atom="http://www.w3.org/2005/Atom"
    xmlns:newznab="http://www.newznab.com/DTD/2010/feeds/attributes/" version="2.0">
    <channel>
        <atom:link href="http://example.com/api?apikey=something&amp;t=search&amp;q=TPB.AFK.2013.360P.H264-20-40" rel="self" type="application/rss+xml"/>
        <title>Your Favorite NZB Indexer</title>
        <description>NZB Feed</description>
        <link>http://example.com/</link>
        <language>en-gb</language>
        <webMaster>webmaster@example.com</webMaster>
        <category/>
        <image>
            <url>http://example.com/images/logo.png</url>
            <title>Example.com</title>
            <link>http://example.com/</link>
            <description>Visit Example.com -</description>
        </image>
        <newznab:response offset="0" total="2"/>
        <item>
            <title>TPB.AFK.2013.360p.H264-20-40</title>
            <guid isPermaLink="true">http://example.com/details/5981bcf2d493a2d811012ade8a9e4082</guid>
            <link>http://example.com/getnzb/5981bcf2d493a2d811012ade8a9e4082.nzb&amp;i=119858</link>
            <comments>http://example.com/details/5981bcf2d493a2d811012ade8a9e4082#comments</comments>
            <pubDate>Mon, 29 Sep 2014 00:54:05 -0400</pubDate>
            <category>Misc > Other</category>
            <description>TPB.AFK.2013.360p.H264-20-40</description>
            <enclosure url="http://example.com/getnzb/5981bcf2d493a2d811012ade8a9e4082.nzb&amp;i=119858" length="440771523" type="application/x-nzb"/>
            <newznab:attr name="category" value="8000"/>
            <newznab:attr name="category" value="8010"/>
            <newznab:attr name="size" value="440771523"/>
            <newznab:attr name="guid" value="5981bcf2d493a2d811012ade8a9e4082"/>
	    	<oznzb:attr name="oz_spam_confirmed" value="no"/>
		<oznzb:attr name="oz_num_passworded_reports" value="0"/>
		<oznzb:attr name="oz_passworded_confirmed" value="no"/>
		<oznzb:attr name="oz_up_votes" value="0"/>
		<oznzb:attr name="oz_down_votes" value="0"/>
		<oznzb:attr name="oz_video_quality_rating" value="0"/>
		<oznzb:attr name="oz_audio_quality_rating" value="0"/>
        </item>
        <item>
            <title>TPB.AFK.2013.360p.H264-20-40</title>
            <guid isPermaLink="true">http://example.com/details/e5d6b9a0b7c266d78b5400f3392c88ee</guid>
            <link>http://example.com/getnzb/e5d6b9a0b7c266d78b5400f3392c88ee.nzb&amp;i=119858</link>
            <comments>http://example.com/details/e5d6b9a0b7c266d78b5400f3392c88ee#comments</comments>
            <pubDate>Mon, 02 Sep 2013 01:24:22 -0400</pubDate>
            <category>Movies > SD</category>
            <description>TPB.AFK.2013.360p.H264-20-40</description>
            <enclosure url="http://example.com/getnzb/e5d6b9a0b7c266d78b5400f3392c88ee.nzb&amp;i=119858&amp;r=5ee3268334a951478294d5fd3d7dd83f" length="406779559" type="application/x-nzb"/>
            <newznab:attr name="category" value="2000"/>
            <newznab:attr name="category" value="2030"/>
            <newznab:attr name="size" value="406779559"/>
            <newznab:attr name="guid" value="e5d6b9a0b7c266d78b5400f3392c88ee"/>
	    	<oznzb:attr name="oz_spam_confirmed" value="no"/>
		<oznzb:attr name="oz_num_passworded_reports" value="0"/>
		<oznzb:attr name="oz_passworded_confirmed" value="no"/>
		<oznzb:attr name="oz_up_votes" value="0"/>
		<oznzb:attr name="oz_down_votes" value="0"/>
		<oznzb:attr name="oz_video_quality_rating" value="0"/>
		<oznzb:attr name="oz_audio_quality_rating" value="0"/>
        </item>
    </channel>
</rss>`)
