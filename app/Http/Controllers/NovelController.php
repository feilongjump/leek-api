<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Symfony\Component\DomCrawler\Crawler;

class NovelController extends Controller
{
    public function search(Request $request)
    {
        $request->validate([
            'keyword' => 'required',
        ], [
            'keyword' => '关键词',
        ]);

        $source = config('novel.source');
        $params = [
            'keyword' => $request->keyword
        ];
        $url = $source . '/search.php';
        $response = Http::asForm()->post($url, $params);

        $crawler = new Crawler($response->body());
        $data = [
            'header' => [],
            'body' => []
        ];
        $crawler->filter('table > tr')->each(function (Crawler $node, $i) use (& $data) {
            [$category, $title, $latestChapter, $author, $updated_at, $status] = explode(' ', $node->text());
            $item = compact('category', 'title', 'latestChapter', 'author', 'updated_at', 'status');

            if ($i === 0) {
                $data['header'] = $item;
            } else {
                $item['link'] = base64_encode($node->filter('a')->link()->getUri());
                $data['body'][] = $item;
            }
        });

        return $data;
    }

    public function show($novel, Request $request)
    {
        $url = base64_decode($novel);

        $response = Http::get($url);

        $crawler = new Crawler($response->body());

        // Book information
        [$info] = $crawler->filter('.info')->each(function (Crawler $node) {
            $info['image'] = $node->filter('img')->image()->getUri();
            $info['title'] = $node->filter('h1')->text();
            [$author, $progress, $category, $updated_at] = $node->filter('.btitle > i')->each(function (Crawler $childrenNode) {
                return $childrenNode->text();
            });

            $introduce = $node->filter('.intro')->text();

            return array_merge(
                $info,
                compact('author', 'progress', 'category', 'updated_at', 'introduce')
            );
        });

        // Book chapters
        $source = config('novel.source');
        [$chapters] = $crawler->filter('#at')->nextAll()->each(function (Crawler $node) use ($source) {
            return $node->filter('tr > td > a')->each(function (Crawler $aNode, $i) use ($source) {
                return [
                    'title' => $aNode->text(),
                    'link' => base64_encode($source . $aNode->attr('href'))
                ];
            });
        });

        return compact('info', 'chapters');
    }
}
